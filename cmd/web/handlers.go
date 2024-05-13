package main

import (
	"net/http"
)

// VirtualTerminal renders terminal.page.gohtml file 
func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["publishable_key"] = app.config.stripe.key
	if err := app.renderTemplate(w, r, "terminal", &templateData{StringMap: stringMap}); err != nil {
		app.errorLog.Println(err)
	}
}

// PaymentSucceeded parses the HTML form data and extracts the relevant payment information,
// and renders the "succeeded" template with the provided data.
//
// Parameters:
//   w (http.ResponseWriter): The HTTP response writer to send the rendered template to the client.
//   r (*http.Request): The HTTP request containing the form data.

func (app *application) PaymentSucceeded(w http.ResponseWriter, r *http.Request) {
	// parsing html form
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	// read form data

	cardHolder := r.Form.Get("cardholder_name") // fields as present in form
	email := r.Form.Get("cardholder_email")
	paymentIntent := r.Form.Get("payment_intent")
	paymentMethod := r.Form.Get("payment_method")
	paymentAmount := r.Form.Get("payment_amount")
	paymentCurrency := r.Form.Get("payment_currency")

	data := make(map[string]interface{})
	data["cardholder"] = cardHolder
	data["email"] = email
	data["pi"] = paymentIntent
	data["pm"] = paymentMethod
	data["pa"] = paymentAmount
	data["pc"] = paymentCurrency

	err = app.renderTemplate(w, r, "succeeded", &templateData{Data: data}) // Data is being accessed in gohtml template
	if err != nil {
		app.errorLog.Println(err)
		return
	}
}
