package main

import (
	"fmt"
	"html"
	"net/http"
	"time"
)

func main() {
	go serverStartup()

	time.Sleep(100 * time.Second)
}

type HandlerFoo struct {
}

func (h HandlerFoo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func serverStartup() {
	handlerFoo := new(HandlerFoo)
	http.Handle("/foo", handlerFoo)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.ListenAndServe(":8080", nil)
}
