package main

import (
	"log"
	"net/http"
)

// handler
func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Snippetbox"))
}

func main() {
    // servemux/router
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)

    log.Print("Starting server on :4000")
    // server
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
