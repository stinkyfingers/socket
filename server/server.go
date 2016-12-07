package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/stinkyfingers/easyrouter"
	"github.com/stinkyfingers/socket/server/db"
	"github.com/stinkyfingers/socket/server/game"
	"github.com/stinkyfingers/socket/server/handlers"
	"golang.org/x/net/websocket"
	"gopkg.in/mgo.v2/bson"
)

var (
	port = flag.String("port", "7000", "server port")
)

func main() {

	// ws
	h := handlers.NewHub()
	go h.Run()

	chub := handlers.NewChatHub()
	go chub.Run()

	var routes = []easyrouter.Route{
		{
			Path:        "/play/{id}",
			Method:      "GET",
			Middlewares: []easyrouter.Middleware{Cors},
			WSHandler: websocket.Handler(func(ws *websocket.Conn) {
				handlers.ServeWS(ws, h)
			}),
		}, {
			Path:        "/chat/{id}",
			Method:      "GET",
			Middlewares: []easyrouter.Middleware{Cors},
			WSHandler: websocket.Handler(func(ws *websocket.Conn) {
				handlers.ChatHandler(ws, chub)
			}),
		}, {
			Path:        "/game/new",
			Method:      "POST",
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleNewGame,
		}, {
			Path:        "/game/add/{id}",
			Method:      "POST",
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleAddPlayer,
		}, {
			Path:        "/game/init/{id}",
			Method:      "GET",
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleStartGame,
		}, {
			Path:        "/game/exit/{id}",
			Method:      "GET",
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleExitGame,
		}, {
			Path:        "/game/update",
			Method:      "PUT",
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleUpdateGame,
		}, {
			Path:        "/game/{id}",
			Method:      "GET",
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleGetGame,
		}, {
			Path:        "/player/reset/{id}",
			Method:      "GET",
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleResetPassword,
		}, {
			Path:        "/player",
			Method:      "PUT",
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleUpdatePlayer,
		}, {
			Path:        "/player",
			Method:      "POST",
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleCreatePlayer,
		}, {
			Path:        "/player",
			Method:      "GET",
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleGetPlayer,
		}, {
			Path:    "/test",
			Method:  "GET",
			Handler: handlers.HandleTestSetup,
		}, {
			Path:    "/import/dealer/{id}",
			Method:  "POST",
			Handler: handlers.HandleImportDealerCards,
		}, {
			Path:    "/import/card/{id}",
			Method:  "POST",
			Handler: handlers.HandleImportCards,
		}, {
			Path:        "/card",
			Method:      "POST",
			Handler:     handlers.HandleCreateCard,
			Middlewares: []easyrouter.Middleware{Cors},
		}, {
			Path:        "/card/{user}",
			Method:      "PUT",
			Handler:     handlers.HandleUpdateCard,
			Middlewares: []easyrouter.Middleware{Cors, SuperUser},
		}, {
			Path:        "/dealer",
			Method:      "POST",
			Handler:     handlers.HandleCreateDealerCard,
			Middlewares: []easyrouter.Middleware{Cors},
		}, {
			Path:        "/dealer/{user}",
			Method:      "PUT",
			Handler:     handlers.HandleUpdateDealerCard,
			Middlewares: []easyrouter.Middleware{Cors, SuperUser},
		}, {
			Path:        "/unreviewed",
			Method:      "GET",
			Handler:     handlers.HandleUnreviewedCards,
			Middlewares: []easyrouter.Middleware{Cors},
		}, {
			Path:    "/export/dealer/{file}",
			Method:  "GET",
			Handler: handlers.HandleExportDealerDeck,
		}, {
			Path:    "/export/cards/{file}",
			Method:  "POST",
			Handler: handlers.HandleExportDeck,
		}, {
			Path:        "/auth",
			Method:      "POST",
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleAuthenticate,
		}, {
			Path:    "/status",
			Method:  "GET",
			Handler: handlers.HandleDefault,
		}, {
			Path:    "/",
			Handler: handlers.HandleDefault,
		},
	}

	err := db.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	s := easyrouter.Server{
		Port:   *port,
		Routes: routes,
		DefaultRoute: easyrouter.Route{
			Path:    "/status",
			Method:  "GET",
			Handler: handlers.HandleDefault,
		},
		Middlewares: []easyrouter.Middleware{Cors},
	}

	s.Run()

}

func Cors(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			return
		}
		fn(w, r)
	}
}

func SuperUser(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.URL.Query().Get("user")
		if !bson.IsObjectIdHex(user) {
			handlers.HttpError{nil, "No user ID", 500, w}.HandleErr()
			return
		}
		p := game.Player{ID: bson.ObjectIdHex(user)}
		err := p.Get()
		if err != nil {
			handlers.HttpError{err, "Error getting user", 500, w}.HandleErr()
			return
		}
		if !p.GrandPoobah {
			handlers.HttpError{nil, "User is not Grand Poobah", 500, w}.HandleErr()
			return
		}
		fn(w, r)
	}
}
