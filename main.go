package main

import (
	"log"
	"net/http"
	"strconv"
	"fmt"
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
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }
    fmt.Fprintf(w, "Snippet with id: %v\n", id)
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
        w.Header().Set("Allow", http.MethodPost)
        w.Header().Set("Content-Type", "text/html")
        http.Error(w, "Method Not Allowed, Use POST", http.StatusMethodNotAllowed)
        return
    })

    log.Print("Starting server on :4000")
    // server
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
