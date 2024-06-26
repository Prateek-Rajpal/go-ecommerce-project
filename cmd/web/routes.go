package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/virtual-terminal", app.VirtualTerminal)
	mux.Get("/charge-once", app.ChargeOnce)
	mux.Post("/payment-succeeded", app.PaymentSucceeded)

	// serving from external directory (static directory)

	fileServer := http.FileServer(http.Dir("./static")) 
	mux.Handle("/static/*", http.StripPrefix("/static",fileServer))

	return mux
}
