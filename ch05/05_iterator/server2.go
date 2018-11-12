package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("tmpl2.html")
	//daysOfWeek := []string{"Mon", "Tue", "Web", "Thu", "Fri", "Sat", "Sun"}
	daysOfWeek := []string{}
	t.Execute(w, daysOfWeek)
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
