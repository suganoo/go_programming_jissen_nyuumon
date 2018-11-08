package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Helloooooo")
}

func world(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "woooorld")
}

func main() {
	server := http.Server {
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
