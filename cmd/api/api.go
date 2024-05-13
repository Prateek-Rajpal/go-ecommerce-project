package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
	stripe struct { // stripe payment gateway details
		secret string
		key    string
	}
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
}

// creating web server
func (app *application) serve() error {

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	app.infoLog.Printf("Starting http server on %s mode on port %d", app.config.env, app.config.port)
	return srv.ListenAndServe()
}

// we make binary out of this package
func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4001, "server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development|production}")
	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime) // for creating info log

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile) // for creating error log

	app := application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
	}

	err := app.serve()

	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
