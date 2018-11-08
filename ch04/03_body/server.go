package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request){
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}
func others(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, r.Method)
	fmt.Fprintln(w, r.Host)
	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.Trailer)
	fmt.Fprintln(w, r.RemoteAddr)
	fmt.Fprintln(w, r.RequestURI)
}

func main() {
	server := http.Server{
		Addr : "127.0.0.1:8080",
	}
	http.HandleFunc("/body", body)
	http.HandleFunc("/others", others)
	server.ListenAndServe()
}
