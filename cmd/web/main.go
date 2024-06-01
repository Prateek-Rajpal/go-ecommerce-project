// this is entrypoint for our frontend application

package main

import (
	"ecommerce_app/internal/driver"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"
const cssVersion = "1"

type config struct {
	port int
	env  string
	api  string // url to call to backend api
	db   struct {
		dsn string // database connection string
	}
	stripe struct { // stripe payment gateway details
		secret string
		key    string
	}
}

type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
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

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development|production}")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to api")
	flag.StringVar(&cfg.db.dsn,"dsn","postgres://postgres:yourpassword@localhost/postgres?sslmode=disable","DSN")
	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime) // for creating info log

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile) // for creating error log
	// database connection
	conn, err := driver.OpenDB(cfg.db.dsn)
	if err!= nil {
        errorLog.Fatal(err) // exit 
    }
	defer conn.Close() // close the connection pool after application stops or exits
	tc := make(map[string]*template.Template)

	app := application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
	}

	err = app.serve()

	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}

}
