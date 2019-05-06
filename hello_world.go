package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

func main() {
	router := mux.NewRouter()
	// Routes consist of a path and a handler function.
	router.HandleFunc("/", RequestHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8001", router))
}
