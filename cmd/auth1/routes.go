package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.HandlerFunc(app.home))

	return mux
}
