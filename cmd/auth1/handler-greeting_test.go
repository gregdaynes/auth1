package main

import (
	"strings"
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
)

func TestGetGreeting(t *testing.T) {
	app := newTestApplication(t)
	_, api := humatest.New(t)
	app.routes(api)

	resp := api.Get("/greeting/world")
	if !strings.Contains(resp.Body.String(), "Hello, world! (1)") {
		t.Fatalf("Unexpected response: %s", resp.Body.String())
	}
}
