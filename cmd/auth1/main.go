package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

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

	cfg, err := config.NewConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		debug:  *debug,
		logger: logger,
		config: &cfg,
	}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
	}

	// logger.Info("starting server", slog.String("addr", *addr))
	l, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}

	host := fmt.Sprint(l.Addr().(*net.TCPAddr))
	x := fmt.Sprintf("Service running on http://%v", host)

	logger.Info(x, slog.String("addr", host))

	log.Fatal(srv.Serve(l))
}
