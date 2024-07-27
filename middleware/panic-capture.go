package middleware

import (
	"github.com/gnulinuxindia/internet-chowkidar/logger/tlog"

	"github.com/go-errors/errors"
	ogenmw "github.com/ogen-go/ogen/middleware"
)

func PanicCapture() ogenmw.Middleware {
	return func(req ogenmw.Request, next ogenmw.Next) (resp ogenmw.Response, respErr error) {
		defer func() {
			if err := recover(); err != nil {
				err := errors.Wrap(err, 1)
				tlog.Exception(req.Context, err)

				resp = ogenmw.Response{}
				respErr = errors.New("internal server error")
			}
		}()

		resp, respErr = next(req)
		return
	}
}
