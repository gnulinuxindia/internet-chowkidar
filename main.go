package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/internal/db"
	"github.com/gnulinuxindia/internet-chowkidar/middleware"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/di"
	"github.com/gnulinuxindia/internet-chowkidar/utils"
	"github.com/rs/cors"

	_ "github.com/gnulinuxindia/internet-chowkidar/migrations" // necessary to register migrations
	"github.com/go-errors/errors"
)

func main() {
	ctx := context.Background()

	// custom logger
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// slog.SetDefault(logger)

	//slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Info("starting server")

	// config
	conf, err := di.InjectConfig()
	if err != nil {
		handleErr(err)
	}

	// auto-migrate
	rawDb, err := di.InjectRawDb()
	if err != nil {
		handleErr(err)
	}

	// migrate to the latest version
	err = db.MigrateUp(rawDb)
	if err != nil {
		handleErr(err)
	}

	// Configure Tracer Provider
	tp, err := di.InjectTracerProvider()
	if err != nil {
		handleErr(err)
	}

	// start http server
	handlers, err := di.InjectHandlers()
	if err != nil {
		slog.Error("error creating handlers", "err", err)
		handleErr(err)
	}

	secHandler, err := di.InjectSecurityHandler()
	if err != nil {
		handleErr(err)
	}

	var server *genapi.Server
	if server, err = genapi.NewServer(
		handlers,
		secHandler,
		genapi.WithTracerProvider(tp),
		genapi.WithMiddleware(
			middleware.PanicCapture(),
			middleware.InfoRec(),
		),
	); err != nil {
		handleErr(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", server)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	httpServer := http.Server{
		ReadHeaderTimeout: time.Second,
		Addr:              fmt.Sprintf("%s:%s", conf.Listen, conf.Port),
		Handler:           c.Handler(mux),
	}

	go func() {
		slog.Info("running server on", "addr", httpServer.Addr)
		err := httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			handleErr(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		slog.Error("shutdown err", "err", err)
		//nolint:gocritic
		os.Exit(1)
	}
}

func handleErr(err error) {
	tperr := errors.Wrap(err, 0)
	fmt.Printf("%s", utils.GetStack(tperr, true))
	panic(err)
}
