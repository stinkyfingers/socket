package main

import (
	"flag"
	"log"
	"net/http"
	"os"

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
			Methods:     []string{"GET"},
			Middlewares: []easyrouter.Middleware{Cors},
			WSHandler: websocket.Handler(func(ws *websocket.Conn) {
				handlers.ServeWS(ws, h)
			}),
		}, {
			Path:        "/chat/{id}",
			Methods:     []string{"GET"},
			Middlewares: []easyrouter.Middleware{Cors},
			WSHandler: websocket.Handler(func(ws *websocket.Conn) {
				handlers.ChatHandler(ws, chub)
			}),
		}, {
			Path:        "/game/new",
			Methods:     []string{"POST"},
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleNewGame,
		}, {
			Path:        "/game/add/{id}",
			Methods:     []string{"POST"},
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleAddPlayer,
		}, {
			Path:        "/game/init/{id}",
			Methods:     []string{"GET"},
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleStartGame,
		}, {
			Path:        "/game/exit/{id}",
			Methods:     []string{"GET"},
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleExitGame,
		}, {
			Path:        "/game/update",
			Methods:     []string{"PUT"},
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleUpdateGame,
		}, {
			Path:        "/game/{id}",
			Methods:     []string{"GET"},
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleGetGame,
		}, {
			Path:        "/player/reset/{id}",
			Methods:     []string{"GET"},
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleResetPassword,
		}, {
			Path:        "/player",
			Methods:     []string{"PUT"},
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleUpdatePlayer,
		}, {
			Path:        "/player",
			Methods:     []string{"POST"},
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleCreatePlayer,
		}, {
			Path:        "/player",
			Methods:     []string{"GET"},
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleGetPlayer,
		}, {
			Path:    "/test",
			Methods: []string{"GET"},
			Handler: handlers.HandleTestSetup,
		}, {
			Path:    "/import/dealer/{id}",
			Methods: []string{"POST"},
			Handler: handlers.HandleImportDealerCards,
		}, {
			Path:    "/import/card/{id}",
			Methods: []string{"POST"},
			Handler: handlers.HandleImportCards,
		}, {
			Path:        "/card",
			Methods:     []string{"POST"},
			Handler:     handlers.HandleCreateCard,
			Middlewares: []easyrouter.Middleware{Cors},
		}, {
			Path:        "/card/{user}",
			Methods:     []string{"PUT"},
			Handler:     handlers.HandleUpdateCard,
			Middlewares: []easyrouter.Middleware{Cors, SuperUser},
		}, {
			Path:        "/dealer",
			Methods:     []string{"POST"},
			Handler:     handlers.HandleCreateDealerCard,
			Middlewares: []easyrouter.Middleware{Cors},
		}, {
			Path:        "/dealer/{user}",
			Methods:     []string{"PUT"},
			Handler:     handlers.HandleUpdateDealerCard,
			Middlewares: []easyrouter.Middleware{Cors, SuperUser},
		}, {
			Path:        "/unreviewed",
			Methods:     []string{"GET"},
			Handler:     handlers.HandleUnreviewedCards,
			Middlewares: []easyrouter.Middleware{Cors},
		}, {
			Path:    "/export/dealer/{file}",
			Methods: []string{"GET"},
			Handler: handlers.HandleExportDealerDeck,
		}, {
			Path:    "/export/cards/{file}",
			Methods: []string{"POST"},
			Handler: handlers.HandleExportDeck,
		}, {
			Path:        "/auth",
			Methods:     []string{"POST"},
			Middlewares: []easyrouter.Middleware{Cors},
			Handler:     handlers.HandleAuthenticate,
		}, {
			Path:    "/status",
			Methods: []string{"GET"},
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
	port_env := os.Getenv("PORT")
	if port_env != "" {
		*port = port_env
	}

	s := easyrouter.Server{
		Port:   *port,
		Routes: routes,
		DefaultRoute: easyrouter.Route{
			Path:    "/status",
			Methods: []string{"GET"},
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
