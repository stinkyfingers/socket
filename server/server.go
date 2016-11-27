package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/stinkyfingers/easyrouter"
	"github.com/stinkyfingers/socket/server/db"
	"github.com/stinkyfingers/socket/server/handlers"
	"golang.org/x/net/websocket"
)

var (
	port = flag.String("port", "7000", "server port")
)

func main() {

	// ws
	h := handlers.NewHub()
	go h.Run()

	var routes = []easyrouter.Route{
		{
			Path:        "/{id}",
			Method:      "GET",
			Middlewares: []easyrouter.Middleware{Cors},
			WSHandler: websocket.Handler(func(ws *websocket.Conn) {
				handlers.ServeWS(ws, h)
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
