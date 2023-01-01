package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// set up flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// set up loggers
	// better to redirect the outputs to a file during execution
	infoLog := log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR \t", log.Ldate|log.Ltime|log.Lshortfile)

	// init application
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// init server
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// start the server at a port 4000
	infoLog.Printf("starting server on %s\n", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
