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

	"github.com/fatih/color"
	"github.com/go-errors/errors"
	_ "github.com/gnulinuxindia/internet-chowkidar/migrations" // necessary to register migrations
)

func main() {
	ctx := context.Background()

	// ─── auto-migrate ────────────────────────────────────────────────
	rawDb, err := di.InjectRawDb()
	if err != nil {
		handleErr(err)
	}

	// migrate to the latest version
	err = db.MigrateUp(rawDb)
	if err != nil {
		handleErr(err)
	}

	// ─── Configure Tracer Provider ────────────────────────────────────────
	tp, err := di.InjectTracerProvider()
	if err != nil {
		handleErr(err)
	}

	// ─── start http server ────────────────────────────────────────────────
	handlers, err := di.InjectHandlers()
	if err != nil {
		handleErr(err)
	}

	var server *genapi.Server
	if server, err = genapi.NewServer(
		handlers,

		// tracer
		genapi.WithTracerProvider(tp),

		// middleware
		genapi.WithMiddleware(
			middleware.PanicCapture(),
			middleware.InfoRec(),
		),
	); err != nil {
		handleErr(err)
	}

	httpServer := http.Server{
		ReadHeaderTimeout: time.Second,
		Addr:              ":9000",
		Handler:           server,
	}

	go func() {
		fmt.Printf("%s running server on %s\n", color.YellowString("[info]"), "localhost:9000")
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
