package main

import (
	"github.com/danielgtaylor/huma/v2"
)

func (app *application) routes(api huma.API) {
	huma.Get(api, "/", app.home)
	huma.Get(api, "/greeting/{name}", app.greeting)
}
