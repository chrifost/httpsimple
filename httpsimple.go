//Simple HTTP server

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func echoString(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello, %q\n", html.EscapeString(r.URL.Path))
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func bot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi\n")
}

func main() {

	// Echo back path
	http.HandleFunc("/", echoString)

	// Do some simple counting
	http.HandleFunc("/count", counter)

	// Echo back bot style
	http.HandleFunc("/hi", bot)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
