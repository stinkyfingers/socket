package main

import (
	"log"
	"net/http"

	"github.com/stinkyfingers/socket/server/db"
	"github.com/stinkyfingers/socket/server/handlers"
	"golang.org/x/net/websocket"
)

func main() {
	err := db.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	// ws
	h := handlers.NewHub()
	go h.Run()

	http.Handle("/", websocket.Handler(func(ws *websocket.Conn) {
		handlers.ServeWS(ws, h)
	}))

	// http
	http.Handle("/game/new", Cors(http.HandlerFunc(handlers.HandleNewGame)))
	http.Handle("/game/add", Cors(http.HandlerFunc(handlers.HandleAddPlayer)))
	http.Handle("/game/init", Cors(http.HandlerFunc(handlers.HandleStartGame)))
	http.Handle("/game/exit", Cors(http.HandlerFunc(handlers.HandleExitGame)))
	http.Handle("/game/update", Cors(http.HandlerFunc(handlers.HandleUpdateGame)))
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
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
