package handler

import (
	"context"
	"time"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/logger/tlog"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/service"
	"github.com/gnulinuxindia/internet-chowkidar/utils"
)

type CounterHandler interface {
	GetCurrentCount(ctx context.Context) (*genapi.Counter, error)
	IncrementCount(ctx context.Context, req genapi.OptIncrement) (*genapi.Counter, error)
}

type counterHandlerImpl struct {
	counterService service.CounterService
}

// TODO: properly implement the two functions
func (h *counterHandlerImpl) GetCurrentCount(ctx context.Context) (*genapi.Counter, error) {
	tlog.Info(ctx, "Testing logs in trace", &map[string]any{
		"meaningOfLife": "42",
	})

	// var a []string
	// fmt.Println(a[3])
	err := utils.Span(ctx, "sleeping zzzz", func(ctx context.Context) error {
		tlog.Info(ctx, "log inside context", nil)
		time.Sleep(1 * time.Second)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &genapi.Counter{
		Count: 0,
	}, nil
}

func (h *counterHandlerImpl) IncrementCount(ctx context.Context, req genapi.OptIncrement) (*genapi.Counter, error) {
	return &genapi.Counter{
		Count: 1,
	}, nil
}
