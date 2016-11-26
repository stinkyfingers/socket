package game

import (
	"errors"
	"math/rand"
	"strconv"

	"github.com/stinkyfingers/socket/server/db"
	"gopkg.in/mgo.v2/bson"
)

type Game struct {
	ID           bson.ObjectId     `bson:"_id" json:"_id"`
	Round        Round             `bson:"round" json:"round"`
	Players      []Player          `bson:"players" json:"players"`
	Initialized  bool              `bson:"initialized" json:"initialized"`
	DealerDeck   []DealerCard      `bson:"dealerDeck" json:"dealerDeck"`
	Deck         []Card            `bson:"deck" json:"deck"`
	FinalScore   map[string][]Play `bson:"finalScore,omitempty" json:"finalScore,omitempty"` // PlayerIDHex to []Vote
	StartedBy    string            `bson:"startedBy" json:"startedBy"`
	Rounds       []Round           `bson:"rounds" json:"rounds"`
	RoundsInGame int               `bson:"roundsInGame" json:"roundsInGame"`
}

type Round struct {
	DealerCards       []DealerCard      `bson:"dealerCards" json:"dealerCards"`
	Plays             map[string]Play   `bson:"plays,omitempty" json:"plays,omitempty"`
	Votes             map[string]Play   `bson:"votes,omitempty" json:"votes,omitempty"`
	Options           []Play            `bson:"options,omitempty" json:"options,omitempty"`
	Score             map[string][]Play `bson:"score,omitempty" json:"score,omitempty"`     //TODO is map[string]Play ok?
	MostRecentResults MostRecentResults `bson:"mostRecentResults" json:"mostRecentResults"` // last round's results
}
type MostRecentResults struct {
	DealerCards []DealerCard    `bson:"dealerCards" json:"dealerCards"`
	Votes       map[string]Play `bson:"votes" json:"votes"`
}

type Card struct {
	Phrase   string        `bson:"phrase" json:"phrase"`
	PlayerID bson.ObjectId `bson:"playerId,omitempty" json:"playerId,omitempty"`
}

type DealerCard struct {
	Phrase string `bson:"phrase" json:"phrase"`
}

type Play struct {
	Player   Player   `bson:"player" json:"player"`
	Card     Card     `bson:"card" json:"card"`
	PlayType PlayType `bson:"playType" json:"playType"`
}

type PlayType string

var cardsInHand = 3
var roundsInGame = 4
var maxPlayers = 10
var maxRounds = 10
var collection = "difference-between"

func (g *Game) Get() error {
	err := db.Session.DB(db.DB).C(collection).FindId(g.ID).One(&g)
	g.Round.Plays = make(map[string]Play)
	g.Round.Votes = make(map[string]Play)
	g.Round.Score = make(map[string][]Play)
	if len(g.Rounds) == roundsInGame {
		return g.TallyScore()
	}
	return err
}

func (g *Game) Update() error {
	return db.Session.DB(db.DB).C(collection).UpdateId(g.ID, g)
}

func (g *Game) Create() error {
	g.ID = bson.NewObjectId()
	if len(g.Players) == 1 {
		g.StartedBy = g.Players[0].ID.Hex()
	}
	return db.Session.DB(db.DB).C(collection).Insert(&g)
}

func (g *Game) AddPlayer(player Player) error {
	err := g.Get()
	if err != nil {
		return err
	}
	if g.Initialized {
		return errors.New("Game has already started")
	}
	g.Players = append(g.Players, player)
	if len(g.Players) == 1 {
		g.StartedBy = g.Players[0].ID.Hex()
	}
	err = db.Session.DB(db.DB).C(collection).UpdateId(g.ID, g)
	return err
}

func (g *Game) Deal() error {
	var err error
	if !g.ID.Valid() {
		return errors.New("Game does not have a valid id")
	}

	total := len(g.Deck) - 1
	if total < len(g.Players)-1 {
		return errors.New("not enough cards in deck to deal cards")
	}

	for p := range g.Players {
		for i := len(g.Players[p].Hand); i < cardsInHand; i++ {
			index := rand.Intn(total)
			total--
			c := g.Deck[index]
			c.PlayerID = g.Players[p].ID
			thisPlayer := g.Players[p]
			thisPlayer.Hand = append(thisPlayer.Hand, c)
			g.Players[p] = thisPlayer
			g.Deck = append(g.Deck[:index], g.Deck[index+1:]...)
		}
	}
	return err
}

func (g *Game) DrawCards() ([]DealerCard, error) {
	var cards []DealerCard
	total := len(g.DealerDeck) - 1
	if total < 2 {
		return nil, errors.New("Not enough cards left")
	}
	for i := 0; i < 2; i++ {
		index := rand.Intn(total)
		cards = append(cards, g.DealerDeck[index])
		g.DealerDeck = append(g.DealerDeck[:index], g.DealerDeck[index+1:]...)
	}
	return cards, nil
}

func (g *Game) Initialize() error {
	if len(g.Deck) == 0 || len(g.DealerDeck) == 0 {
		return errors.New("Starting game requires a deck & dealer deck")
	}
	if len(g.Players) < 2 {
		return errors.New("Starting game requires two or more players")
	}
	if g.RoundsInGame == 0 {
		g.RoundsInGame = roundsInGame
	}
	if len(g.Players) > maxPlayers {
		return errors.New("Max number of players is " + strconv.Itoa(maxPlayers))
	}
	if g.RoundsInGame > maxRounds {
		return errors.New("Max number of rounds is " + strconv.Itoa(maxRounds))
	}

	g.Initialized = true

	//deal
	err := g.Deal()
	if err != nil {
		return err
	}
	// 1st round
	cards, err := g.DrawCards()
	if err != nil {
		return err
	}
	r := Round{
		DealerCards: cards,
		// Previous:    nil,
		Plays: make(map[string]Play),
		Votes: make(map[string]Play),
		Score: make(map[string][]Play),
	}
	g.Round = r

	err = db.Session.DB(db.DB).C(collection).UpdateId(g.ID, g)
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) DrawNewCard() Card {
	total := len(g.Deck) - 1
	index := rand.Intn(total)
	newCard := g.Deck[index]
	g.Deck = append(g.Deck[:index], g.Deck[index+1:]...)
	return newCard
}

func (g *Game) ReplacePlayerCard(p Play) error {
	var replaced bool
	for pid, player := range g.Players {
		if player.ID != p.Player.ID {
			continue
		}
		for i, c := range g.Players[pid].Hand {
			if c.Phrase == p.Card.Phrase {
				newCard := g.DrawNewCard()
				newCard.PlayerID = g.Players[pid].ID
				g.Players[pid].Hand[i] = newCard
				replaced = true
			}
		}
	}
	if !replaced {
		return errors.New("card not found/replaced")
	}
	return nil
}

func (g *Game) UpdatePlays() error {
	for _, play := range g.Round.Plays {
		// assign other player fields
		for i, player := range g.Players {
			if player.ID == play.Player.ID {
				play.Player = player
				// remove played cards from players
				for j, playerCard := range g.Players[i].Hand {
					if playerCard.Phrase == play.Card.Phrase {
						g.Players[i].Hand = append(g.Players[i].Hand[:j], g.Players[i].Hand[j+1:]...)
					}
				}
			}
		}
		// assign to options
		g.Round.Options = append(g.Round.Options, play)
	}

	// nullify plays
	g.Round.Plays = make(map[string]Play)
	return nil
}

func (g *Game) UpdateVotes() error {
	if g.Round.Score == nil {
		g.Round.Score = make(map[string][]Play)
	}
	for _, play := range g.Round.Votes {
		// assign other player fields
		for _, player := range g.Players {
			if player.ID == play.Player.ID {
				play.Player = player
			}
		}
	}

	// score
	for _, play := range g.Round.Votes {
		g.Round.Score[play.Card.PlayerID.Hex()] = append(g.Round.Score[play.Card.PlayerID.Hex()], play)
	}

	// Next Round
	newDealerCards, err := g.DrawCards()
	if err != nil {
		return err
	}

	// lastRound := g.Round
	g.Rounds = append(g.Rounds, g.Round)
	r := Round{
		DealerCards: newDealerCards,
		Plays:       make(map[string]Play),
		Votes:       make(map[string]Play),
		Score:       make(map[string][]Play),
		MostRecentResults: MostRecentResults{
			Votes:       g.Round.Votes,
			DealerCards: g.Round.DealerCards,
		},
	}
	g.Round = r

	// Check for game end
	if len(g.Rounds) == roundsInGame {
		return g.TallyScore()
	}

	err = g.Deal()
	if err != nil {
		return err
	}

	return g.Update()
}

// TallyScore traverses a game's rounds, appending the scores to a [playerID][]Vote map
func (g *Game) TallyScore() error {
	g.FinalScore = make(map[string][]Play)
	for _, round := range g.Rounds {
		for playerID, score := range round.Score {
			g.FinalScore[playerID] = append(g.FinalScore[playerID], score...)
		}
	}

	g.Round = Round{}

	err := g.Update()
	return err
}
