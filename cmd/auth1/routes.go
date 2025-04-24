package main

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	api := humago.New(mux, huma.DefaultConfig("My API", "1.0.0"))

	huma.Get(api, "/greeting/{name}", app.greeting)

	return mux
}
