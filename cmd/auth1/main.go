package main

import (
	"database/sql"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/gregdaynes/auth1/internal/config"
	migrate "github.com/gregdaynes/auth1/internal/migration"

	_ "modernc.org/sqlite"
)

type application struct {
	debug  bool
	logger *slog.Logger
	config *config.Config
	db     *sql.DB
}

func main() {
	name := flag.String("name", "Auth 1", "Application name")
	addr := flag.String("addr", "127.0.0.1:3000", "HTTP network address")
	debug := flag.Bool("debug", false, "Enable debug mode")
	dsn := flag.String("dsn", "file:auth1.sqlite3", "Path for Sqlite database")
	flag.Parse()

	slogHandlerOptions := slog.HandlerOptions{}
	slogHandlerOptions.Level = slog.LevelInfo

	if *debug {
		slogHandlerOptions.AddSource = true
		slogHandlerOptions.Level = slog.LevelDebug
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slogHandlerOptions))

	cfg, err := config.NewConfiguration(config.Config{
		Name:  *name,
		Addr:  *addr,
		Debug: *debug,
		DbDsn: *dsn,
	})
	if err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open("sqlite", *dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	migrate.Migrate(conn, "./schema.sql")

	app := &application{
		debug:  *debug,
		logger: logger,
		config: &cfg,
		db:     conn,
	}

	router := http.NewServeMux()
	api := humago.New(router, huma.DefaultConfig("My API", "1.0.0"))
	app.routes(api)

	logger.Info("starting server", "addr", *addr)
	http.ListenAndServe(*addr, router)
}
