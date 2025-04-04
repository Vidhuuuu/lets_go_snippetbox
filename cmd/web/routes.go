package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/snippet/view", app.snippetView)

    mux.HandleFunc("POST /snippet/create", app.snippetCreate)
    mux.HandleFunc("/snippet/create", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Allow", http.MethodPost)
        w.Header().Set("Content-Type", "text/html")
		app.clientError(w, http.StatusMethodNotAllowed)
        return
    })

	return mux
}
