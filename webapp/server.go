package main

import (
	"net/http"
)

func init() {

	http.Handle("/static", http.StripPrefix("/static", http.FileServer(http.Dir("./build/"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./build/index.html")
	})

}
