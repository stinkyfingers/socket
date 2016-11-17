package game

import (
	"errors"
	"math/rand"

	"github.com/stinkyfingers/socket/server/db"
	"gopkg.in/mgo.v2/bson"
)

type Game struct {
	ID          bson.ObjectId     `bson:"_id" json:"_id"`
	Round       Round             `bson:"round" json:"round"`
	Players     []Player          `bson:"players" json:"players"`
	Initialized bool              `bson:"initialized" json:"initialized"`
	DealerDeck  []DealerCard      `bson:"dealerDeck" json:"dealerDeck"`
	Deck        []Card            `bson:"deck" json:"deck"`
	FinalScore  map[string][]Play `bson:"finalScore" json:"finalScore"` // PlayerIDHex to []Vote
}

type Round struct {
	DealerCards []DealerCard      `bson:"dealerCards" json:"dealerCards"`
	Plays       map[string]Play   `bson:"plays,omitempty" json:"plays,omitempty"`
	Votes       map[string]Play   `bson:"votes,omitempty" json:"votes,omitempty"`
	Options     []Play            `bson:"options,omitempty" json:"options,omitempty"`
	Score       map[string][]Play `bson:"score,omitempty" json:"score,omitempty"` //TODO is map[string]Play ok?
	Previous    *Round            `bson:"previous" json:"-"`
}

type Card struct {
	Phrase string `bson:"phrase" json:"phrase"`
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
var roundsInGame = 3
var collection = "difference-between"

func (g *Game) Get() error {
	err := db.Session.DB(db.DB).C(collection).FindId(g.ID).One(&g)
	g.Round.Plays = make(map[string]Play)
	g.Round.Votes = make(map[string]Play)
	g.Round.Score = make(map[string][]Play)
	return err
}

func (g *Game) Update() error {
	return db.Session.DB(db.DB).C(collection).UpdateId(g.ID, g)
}

func (g *Game) Create() error {
	g.ID = bson.NewObjectId()
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
	if len(g.Players) < 1 {
		return errors.New("Starting game requires two or more players") //TODO - make 2
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
		Previous:    nil,
		Plays:       make(map[string]Play),
		Votes:       make(map[string]Play),
		Score:       make(map[string][]Play),
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
	for _, play := range g.Round.Votes {
		g.Round.Score[play.Player.ID.Hex()] = append(g.Round.Score[play.Player.ID.Hex()], play)
	}

	// Check for game end
	rounds := g.GetRounds()
	if len(rounds) == roundsInGame {
		return g.TallyScore(rounds)
	}

	// Next Round
	newDealerCards, err := g.DrawCards()
	if err != nil {
		return err
	}
	lastRound := g.Round
	r := Round{
		DealerCards: newDealerCards,
		Previous:    &lastRound,
		Plays:       make(map[string]Play),
		Votes:       make(map[string]Play),
		Score:       make(map[string][]Play),
	}
	g.Round = r
	err = g.Deal()
	if err != nil {
		return err
	}
	return g.Update()
}

// GetRounds traverses a game's rounds into an array and returns it
func (g *Game) GetRounds() []Round {
	var rounds []Round
	currentRound := g.Round
	for {
		rounds = append(rounds, currentRound)
		if currentRound.Previous == nil {
			break
		}
		prev := currentRound.Previous
		currentRound = *prev
	}
	return rounds
}

// TallyScore traverses a game's rounds, appending the votes (plays) to a [playerID][]Vote map
func (g *Game) TallyScore(rounds []Round) error {
	g.FinalScore = make(map[string][]Play)
	for _, round := range rounds {
		for _, vote := range round.Votes {
			g.FinalScore[vote.Player.ID.Hex()] = append(g.FinalScore[vote.Player.ID.Hex()], vote)
		}
	}
	return nil
}