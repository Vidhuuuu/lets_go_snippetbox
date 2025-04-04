package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    tmplFiles := []string{
        "./ui/html/base.tmpl",
        "./ui/html/pages/home.tmpl",
        "./ui/html/partials/nav.tmpl",
    }

    ts, err := template.ParseFiles(tmplFiles...)
    if err != nil {
        app.errorLog.Print(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        app.errorLog.Print(err.Error())
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }
    fmt.Fprintf(w, "Snippet with id: %v\n", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Create snippets"))
}
