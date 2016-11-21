package main

import (
	"log"
	"net/http"

	"github.com/stinkyfingers/socket/server/db"
	"github.com/stinkyfingers/socket/server/game"
	"github.com/stinkyfingers/socket/server/handlers"
	"golang.org/x/net/websocket"
)

func main() {
	err := db.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	handlers.Clients = make(map[string][]handlers.Client)
	handlers.Games = make(map[string]game.Game)

	http.Handle("/", websocket.Handler(handlers.Game))

	http.Handle("/game/new", Cors(http.HandlerFunc(handlers.HandleNewGame)))
	http.Handle("/game/add", Cors(http.HandlerFunc(handlers.HandleAddPlayer)))
	http.Handle("/game/init", Cors(http.HandlerFunc(handlers.HandleStartGame)))
	http.Handle("/game/exit", Cors(http.HandlerFunc(handlers.HandleExitGame)))
	http.Handle("/game", Cors(http.HandlerFunc(handlers.HandleGetGame)))

	http.Handle("/player/reset", Cors(http.HandlerFunc(handlers.HandleResetPassword)))
	http.Handle("/player/update", Cors(http.HandlerFunc(handlers.HandleUpdatePlayer)))
	http.Handle("/player/create", Cors(http.HandlerFunc(handlers.HandleCreatePlayer)))
	http.Handle("/player", Cors(http.HandlerFunc(handlers.HandleGetPlayer)))

	http.HandleFunc("/test", handlers.HandleTestSetup)
	http.Handle("/auth", Cors(http.HandlerFunc(handlers.HandleAuthenticate)))
	http.Handle("/status", http.HandlerFunc(handlers.HandleDefault))
	log.Fatal(http.ListenAndServe(":7000", nil))
}

func Cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

// type Client struct {
// 	ws *websocket.Conn
// 	IP string
// }

// var clients map[string][]Client
// var games map[string]game.Game

// func handler(ws *websocket.Conn) {
// 	id := ws.Request().URL.Query().Get("id")

// 	var g game.Game
// 	g.ID = bson.ObjectIdHex(id)
// 	err := g.Get()
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}
// 	games[g.ID.Hex()] = g

// 	// client return
// 	client := Client{
// 		ws,
// 		ws.Request().RemoteAddr,
// 	}
// 	clients[id] = append(clients[id], client)

// 	for {

// 		for _, c := range clients[id] {
// 			err = websocket.JSON.Send(c.ws, games[id])
// 			if err != nil {
// 				log.Print("WS client connection error: ", err)
// 				// break
// 			}
// 		}

// 		var p game.Play
// 		err = websocket.JSON.Receive(ws, &p)
// 		if err != nil {
// 			log.Print(err)
// 			break //TODO - handle errors in WS
// 		}

// 		switch p.PlayType {
// 		case "play":
// 			games[id].Round.Plays[p.Player.ID.Hex()] = p // TODO - switch to id.hex
// 		case "vote":
// 			games[id].Round.Votes[p.Player.ID.Hex()] = p
// 		default:
// 			log.Print("type not supported")
// 		}

// 		ch := make(chan game.Game)
// 		go func() {
// 			if p.PlayType == "play" && len(games[id].Round.Plays) > 1 { //TODO -len equal to players
// 				ga := games[id]
// 				(&ga).UpdatePlays()
// 				ch <- ga
// 			} else if p.PlayType == "vote" && len(games[id].Round.Votes) > 1 {
// 				ga := games[id]
// 				(&ga).UpdateVotes()
// 				ch <- ga
// 			} else {
// 				ch <- games[id]
// 			}
// 		}()

// 		ga := <-ch
// 		games[id] = ga
// 	}
// }
