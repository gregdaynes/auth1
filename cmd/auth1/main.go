package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/gregdaynes/auth1/internal/config"
)

type application struct {
	debug  bool
	logger *slog.Logger
	config *config.Config
}

func main() {
	name := flag.String("name", "Auth 1", "Application name")
	addr := flag.String("addr", "127.0.0.1:3000", "HTTP network address")
	debug := flag.Bool("debug", false, "Enable debug mode")
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
	})
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		debug:  *debug,
		logger: logger,
		config: &cfg,
	}

	router := http.NewServeMux()
	api := humago.New(router, huma.DefaultConfig("My API", "1.0.0"))
	app.routes(api)

	logger.Info("starting server", "addr", *addr)
	http.ListenAndServe(*addr, router)
}
