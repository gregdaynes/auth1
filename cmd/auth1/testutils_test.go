package main

import (
	"io"
	"log/slog"
	"testing"

	"github.com/gregdaynes/auth1/internal/config"
)

func newTestApplication(t *testing.T) *application {
	return &application{
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
		config: &config.Config{
			Name:  "Test 1",
			Addr:  ":6666",
			Debug: false,
		},
	}
}
