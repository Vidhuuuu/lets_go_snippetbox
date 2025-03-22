package main

import (
	"log"
	"net/http"
)

// handler
func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("View snippets"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Create snippets"))
}

func main() {
    // servemux/router
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view", snippetView)

    mux.HandleFunc("POST /snippet/create", snippetCreate)
    mux.HandleFunc("/snippet/create", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Allow", "POST")
        w.WriteHeader(405)
        w.Write([]byte("Method Not Allowed, User POST"))
    })

    log.Print("Starting server on :4000")
    // server
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
