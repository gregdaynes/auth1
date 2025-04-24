package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gregdaynes/auth1/internal/config"
)

type application struct {
	debug  bool
	logger *slog.Logger
	config *config.Config
}

func main() {
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

	router := app.routes()

	addrParts := strings.Split(*addr, ":")
	port, err := strconv.Atoi(addrParts[1])

	app.config.Port = int(port)

	http.ListenAndServe(*addr, router)
}
