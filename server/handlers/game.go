package handlers

import (
	"encoding/json"
	// "log"
	"net/http"

	"github.com/stinkyfingers/socket/server/game"
	// "golang.org/x/net/websocket"
	"gopkg.in/mgo.v2/bson"
)

// type Client struct {
// 	ws *websocket.Conn
// 	IP string
// }

// type Pool struct {
// 	// Clients map[string][]Client
// 	ClientGroups map[string]ClientGroup
// }

// type ClientGroup struct {
// 	Clients []*Client
// 	Game    game.Game
// }

// func NewPool() *Pool {
// 	return &Pool{
// 		// Clients: make(map[string][]Client),
// 		ClientGroups: make(map[string]ClientGroup),
// 	}
// }

// func (p *Pool) GetClients(id string) []*Client {
// 	return p.ClientGroups[id].Clients
// }

// func (p *Pool) RemoveClient(c Client, id string) {
// 	index := -1
// 	for i := range p.ClientGroups[id].Clients {
// 		if p.ClientGroups[id].Clients[i].IP == c.IP {
// 			index = i
// 		}
// 	}
// 	if index > -1 {
// 		cg := p.ClientGroups[id]
// 		cg.Clients = append(p.ClientGroups[id].Clients[:index], p.ClientGroups[id].Clients[index+1:]...)

// 		p.ClientGroups[id] = cg
// 	}
// }

// func (p *Pool) AddClient(c Client, id string) {
// 	cg := p.ClientGroups[id]
// 	cg.Clients = append(cg.Clients, &c)
// 	p.ClientGroups[id] = cg

// }

// func (p *Pool) AddGame(g game.Game) {
// 	cg := p.ClientGroups[g.ID.Hex()]
// 	cg.Game = g
// 	p.ClientGroups[g.ID.Hex()] = cg
// }

// // var Clients map[string][]Client
// // var Games map[string]game.Game

// // Server needs to:
// // 1. Send each client Lead Cards and their Player Hand
// // 2. Receive each client's play
// // 3. Send each client Played Cards
// // 4. Receive each client's vote

// func Game(ws *websocket.Conn, p *Pool) {
// 	id := ws.Request().URL.Query().Get("id")
// 	log.Print("ID ", id)
// 	if id == "" {
// 		return
// 	}

// 	var g game.Game
// 	g.ID = bson.ObjectIdHex(id)
// 	err := g.Get()
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	// Games[g.ID.Hex()] = g
// 	p.AddGame(g)

// 	// client return
// 	client := Client{
// 		ws,
// 		ws.Request().RemoteAddr,
// 	}
// 	// Clients[id] = append(Clients[id], client)
// 	p.AddClient(client, id)
// 	p.handle(id, ws)
// }

// func (pool *Pool) handle(id string, ws *websocket.Conn) {
// 	sendChan := make(chan game.Game, 1)
// 	done := make(chan int)
// 	go pool.Write(id, sendChan, done)

// 	sendChan <- pool.ClientGroups[id].Game

// 	for {

// 		log.Print("WAIT")
// 		<-done
// 		log.Print("READY")

// 		var p game.Play
// 		err := websocket.JSON.Receive(ws, &p)
// 		if err != nil {
// 			log.Print(err)
// 			break //TODO - handle errors in WS
// 		}

// 		switch p.PlayType {
// 		case "play":
// 			pool.ClientGroups[id].Game.Round.Plays[p.Player.ID.Hex()] = p
// 		case "vote":
// 			pool.ClientGroups[id].Game.Round.Votes[p.Player.ID.Hex()] = p
// 		default:
// 			log.Print("type not supported")
// 		}
// 		log.Print("P ", p)

// 		ch := make(chan game.Game)

// 		go func() {
// 			if p.PlayType == "play" && len(pool.ClientGroups[id].Game.Round.Plays) > 1 { //TODO -len equal to players
// 				ga := pool.ClientGroups[id].Game
// 				(&ga).UpdatePlays()
// 				ch <- ga
// 			} else if p.PlayType == "vote" && len(pool.ClientGroups[id].Game.Round.Votes) > 1 {
// 				ga := pool.ClientGroups[id].Game
// 				(&ga).UpdateVotes()
// 				ch <- ga
// 			} else {
// 				ch <- pool.ClientGroups[id].Game
// 			}
// 		}()

// 		cg := pool.ClientGroups[id]
// 		cg.Game = <-ch
// 		pool.ClientGroups[id] = cg

// 		sendChan <- pool.ClientGroups[id].Game

// 	}
// }

// func (p *Pool) Write(id string, ch chan game.Game, done chan int) {
// 	for {
// 		select {
// 		case g := <-ch:

// 			for _, c := range p.ClientGroups[id].Clients {
// 				log.Print("C", *c)
// 				c.ws.SetWriteDeadline(time.Now().Add(time.Second / 3))
// 				err := websocket.JSON.Send(c.ws, g)
// 				if err != nil {
// 					log.Print("WS client connection error: ", err)
// 					// remove conn
// 					p.RemoveClient(*c, id)

// 				}
// 			}
// 			log.Print("DONE")
// 			done <- 1
// 		}
// 	}
// }

// for _, cli := range p.Clients[id] {
// 	// send setup
// 	sendChan := make(chan game.Game)
// 	// go p.writeHandler(sendChan)
// 	go write(cli, sendChan)
// 	sendChan <- g

// 	for {

// 		// sendChan <- Games[id]

// 		// for i, c := range Clients[id] {
// 		// 	log.Print("CLIENT ", c.IP, ws)
// 		// 	err = websocket.JSON.Send(c.ws, Games[id])
// 		// 	if err != nil {
// 		// 		log.Print("WS client connection error: ", err)
// 		// 		Clients[id] = append(Clients[id][:i], Clients[id][i+1:]...) //remove dead client
// 		// 		// break
// 		// 	}
// 		// }

// 		var p game.Play
// 		err = websocket.JSON.Receive(cli.ws, &p)
// 		if err != nil {
// 			log.Print(err)
// 			break //TODO - handle errors in WS
// 		}

// 		switch p.PlayType {
// 		case "play":
// 			Games[id].Round.Plays[p.Player.ID.Hex()] = p // TODO - switch to id.hex
// 		case "vote":
// 			Games[id].Round.Votes[p.Player.ID.Hex()] = p
// 		default:
// 			log.Print("type not supported")
// 		}
// 		log.Print("P ", p)
// 		ch := make(chan game.Game)
// 		go func() {
// 			log.Print(len(Games[id].Round.Plays))
// 			if p.PlayType == "play" && len(Games[id].Round.Plays) > 1 { //TODO -len equal to players
// 				ga := Games[id]
// 				(&ga).UpdatePlays()
// 				ch <- ga
// 			} else if p.PlayType == "vote" && len(Games[id].Round.Votes) > 1 {
// 				ga := Games[id]
// 				(&ga).UpdateVotes()
// 				ch <- ga
// 			} else {
// 				ch <- Games[id]
// 			}
// 		}()
// 		Games[id] = <-ch
// 		sendChan <- Games[id]
// 	}
// }
// }

// func write(cli Client, ch chan game.Game) {
// 	for {
// 		select {
// 		case g := <-ch:
// 			log.Print("sending to ", cli)
// 			err := websocket.JSON.Send(cli.ws, g)
// 			if err != nil {
// 				// p.RemoveClient(c, g.ID.Hex())
// 			}

// 		}
// 	}
// }

// func (p *Pool) writeHandler(ch chan game.Game) {
// 	for {
// 		select {
// 		case g := <-ch:
// 			for _, c := range p.Clients[g.ID.Hex()] {
// 				log.Print("sending to ", c)
// 				err := websocket.JSON.Send(c.ws, g)
// 				if err != nil {
// 					p.RemoveClient(c, g.ID.Hex())
// 				}

// 			}
// 		}
// 	}
// }

// func (p *Pool) readHandler(ws *websocket.Conn, id string) chan game.Game {
// 	var play game.Play
// 	err := websocket.JSON.Receive(ws, &p)
// 	if err != nil {
// 		log.Print(err)
// 		return nil
// 	}

// 	switch play.PlayType {
// 	case "play":
// 		Games[id].Round.Plays[play.Player.ID.Hex()] = play // TODO - switch to id.hex
// 	case "vote":
// 		Games[id].Round.Votes[play.Player.ID.Hex()] = play
// 	default:
// 		log.Print("type not supported")
// 	}
// 	log.Print("PLAY ", play)
// 	ch := make(chan game.Game)
// 	go func() {
// 		log.Print(len(Games[id].Round.Plays))
// 		if play.PlayType == "play" && len(Games[id].Round.Plays) > 1 { //TODO -len equal to players
// 			ga := Games[id]
// 			(&ga).UpdatePlays()
// 			ch <- ga
// 		} else if play.PlayType == "vote" && len(Games[id].Round.Votes) > 1 {
// 			ga := Games[id]
// 			(&ga).UpdateVotes()
// 			ch <- ga
// 		} else {
// 			ch <- Games[id]
// 		}
// 	}()
// 	return ch
// 	// Games[id] = <-ch
// 	// sendChan <- Games[id]
// }

func HandleNewGame(w http.ResponseWriter, r *http.Request) {
	var player game.Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		HttpError{err, "Player required to start game", 500, w}.HandleErr()
		return
	}
	deck, err := game.GetDeck()
	if err != nil {
		HttpError{err, "error getting deck", 500, w}.HandleErr()
		return
	}
	dealerDeck, err := game.GetDealerDeck()
	if err != nil {
		HttpError{err, "error getting dealer deck", 500, w}.HandleErr()
		return
	}
	g := game.Game{
		Deck:       deck.Cards,
		DealerDeck: dealerDeck.Cards,
		Players:    []game.Player{player},
	}
	err = g.Create()
	if err != nil {
		HttpError{err, "error creating game", 500, w}.HandleErr()
		return
	}
	sendJson(w, g)
}

func HandleAddPlayer(w http.ResponseWriter, r *http.Request) {
	var player game.Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		HttpError{err, "Player required to join game", 500, w}.HandleErr()
		return
	}
	id := r.URL.Query().Get("id")
	if !bson.IsObjectIdHex(id) {
		HttpError{err, "error parsing game id", 500, w}.HandleErr()
		return
	}
	g := game.Game{
		ID: bson.ObjectIdHex(id),
	}
	err = g.AddPlayer(player)
	if err != nil {
		HttpError{err, "error adding player", 500, w}.HandleErr()
		return
	}
	sendJson(w, g)
}

func HandleStartGame(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if !bson.IsObjectIdHex(id) {
		HttpError{nil, "error parsing game id", 500, w}.HandleErr()
		return
	}
	g := game.Game{
		ID: bson.ObjectIdHex(id),
	}
	err := g.Get()
	if err != nil {
		HttpError{err, "error retrieving game", 500, w}.HandleErr()
		return
	}
	err = g.Initialize()
	if err != nil {
		HttpError{err, "error initializing game", 500, w}.HandleErr()
		return
	}
	sendJson(w, g)
}

func HandleGetGame(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if !bson.IsObjectIdHex(id) {
		HttpError{nil, "error parsing game id", 500, w}.HandleErr()
		return
	}
	g := game.Game{
		ID: bson.ObjectIdHex(id),
	}
	err := g.Get()
	if err != nil {
		HttpError{err, "error retrieving game", 500, w}.HandleErr()
		return
	}
	sendJson(w, g)
}

func HandleExitGame(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if !bson.IsObjectIdHex(id) {
		HttpError{nil, "error parsing game id", 500, w}.HandleErr()
		return
	}
	g := game.Game{
		ID: bson.ObjectIdHex(id),
	}
	err := g.Get()
	if err != nil {
		HttpError{err, "error retrieving game", 500, w}.HandleErr()
		return
	}
	g.Initialized = false
	err = g.Update()
	if err != nil {
		HttpError{err, "error updating game", 500, w}.HandleErr()
		return
	}
	sendJson(w, g)
}
