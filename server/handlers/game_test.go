package handlers

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGame(t *testing.T) {
	s := server()
	// p := NewPool()

	// handler := http.Handler(websocket.Handler(func(ws *websocket.Conn) { Game(ws, p) }))
	uri := s.URL + "/"
	t.Log(s.URL)
	t.Log(uri)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Error(err)
	}
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)

}

func server() *httptest.Server {
	p := NewPool()
	http.Handle("/", websocket.Handler(func(ws *websocket.Conn) {
		log.Print("HERE")
		Game(ws, p)
	}))
	http.Handle("/foo", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("FOOOOO"))
	}))
	server := httptest.NewServer(nil)
	addr := server.Listener.Addr().String()
	log.Print("Test server running at ", addr)
	return server
}
