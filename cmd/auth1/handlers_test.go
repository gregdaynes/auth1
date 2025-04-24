package main

import (
	"strings"
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
)

func TestGetHome(t *testing.T) {
	app := newTestApplication(t)
	_, api := humatest.New(t)
	app.routes(api)

	resp := api.Get("/")
	if !strings.Contains(resp.Body.String(), "Hello world!") {
		t.Fatalf("Unexpected response: %s", resp.Body.String())
	}
}
