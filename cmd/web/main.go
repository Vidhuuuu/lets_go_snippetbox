package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"github.com/Vidhuuuu/lets_go_snippetbox/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	infoLog *log.Logger
	errorLog *log.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	dsn := flag.String("dsn", "snippetuser:passpass@/snippetbox?parseTime=true", "MariaDB data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate | log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
	    errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application {
		infoLog: infoLog,
		errorLog: errorLog,
		snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

    infoLog.Printf("Starting server on %v\n", *addr)
    err = srv.ListenAndServe()
    errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
	    return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
