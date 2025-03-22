package main

import (
	"log"
	"net/http"
)

func main() {
    mux := http.NewServeMux()

    fileServer := http.FileServer(http.Dir("./ui/static/"))

    mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

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
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
