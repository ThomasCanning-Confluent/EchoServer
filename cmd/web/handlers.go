package main

import (
	"fmt"
	"net/http"
)

/*
Each page has a handler that takes in a http writer and a request path. Checks can be done on the request path to make sure requested path is correct, e.g. for subtree paths restricting their catch-all behaviour.
The rest of the function then dictates what happens on that page, with the http writer being used to write to the page.
*/

// echoHandler is a simple handler that echoes the request headers back to the user
func (app *application) echoHandler(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
