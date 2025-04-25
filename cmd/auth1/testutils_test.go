package main

import (
	"database/sql"
	"io"
	"log/slog"
	"testing"

	"github.com/gregdaynes/auth1/internal/config"
	migrate "github.com/gregdaynes/auth1/internal/migration"
)

func newTestApplication(t *testing.T) *application {
	dsn := "file:../../test.db?mode=memory"
	conn, err := sql.Open("sqlite", dsn)
	if err != nil {
		t.Fatal(err)
	}
	migrate.Migrate(conn, "../../schema.sql")

	return &application{
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
		config: &config.Config{
			Name:  "Test 1",
			Addr:  ":6666",
			Debug: false,
			DbDsn: dsn,
		},
		db: conn,
	}
}
