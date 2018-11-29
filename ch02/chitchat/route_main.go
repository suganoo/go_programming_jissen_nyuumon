package main

import (
	"fmt"
	"./data"
	"net/http"
)

func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		error_message(writer, request, "Cannot get threads")
	} else {
		_, err := session(writer, request)
		fmt.Println("index - session err:")
		fmt.Println(err)
		if err != nil {
			fmt.Println("public")
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			fmt.Println("private")
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
