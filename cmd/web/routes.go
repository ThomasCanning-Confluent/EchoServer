package main

import (
	"github.com/justinas/alice"
	"net/http"
)

// routes returns a http.Handler which routes requests to the correct handler
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.echoHandler)

	// Create a middleware chain containing of standard middleware
	// which will be used for every request our application receives
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// Return the standard middleware chain followed by the servemux.
	return standard.Then(mux)
}
