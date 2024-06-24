package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// holds application wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	//Flags that can be adjusted when run from terminal
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	//Custom logging functions
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//Application holds information used throughout app, such as custom errors and a slice of snippets
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	//A http server that listens on the specified address and uses the routes defined in the routes.go file
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
