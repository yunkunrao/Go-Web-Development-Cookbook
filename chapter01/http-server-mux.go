package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {
	// This allocates and returns a new HTTP request multiplexer (ServeMux), which
	// matches the URL of each incoming request against a list of registered patterns
	// and calls the handler for the pattern that most closely matches the URL.
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	// wraps our server with a .gzip handler to compress all responses in a .gzip format.
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, handlers.CompressHandler(mux))
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
