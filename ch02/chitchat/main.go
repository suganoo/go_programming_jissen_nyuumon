package main

import (
	"net/http"
)

func main() {
	p("ChitChat", version(), "started at", config.Address)

	mux := http.NewServerMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)
}
