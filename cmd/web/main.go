package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog *log.Logger
	errorLog *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate | log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile)

	app := &application {
		infoLog: infoLog,
		errorLog: errorLog,
	}

    mux := http.NewServeMux()

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/snippet/view", app.snippetView)

    mux.HandleFunc("POST /snippet/create", app.snippetCreate)
    mux.HandleFunc("/snippet/create", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Allow", http.MethodPost)
        w.Header().Set("Content-Type", "text/html")
        http.Error(w, "Method Not Allowed, Use POST", http.StatusMethodNotAllowed)
        return
    })

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

    infoLog.Printf("Starting server on %v\n", *addr)
    err := srv.ListenAndServe()
    errorLog.Fatal(err)
}
