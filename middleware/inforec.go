package middleware

import (
	"context"
	"encoding/json"

	"github.com/gnulinuxindia/internet-chowkidar/logger"
	"github.com/gnulinuxindia/internet-chowkidar/logger/tlog"

	ogenmw "github.com/ogen-go/ogen/middleware"
)

func InfoRec() ogenmw.Middleware {
	return func(req ogenmw.Request, next ogenmw.Next) (ogenmw.Response, error) {
		ctx := req.Context
		ctx = context.WithValue(ctx, logger.Host, req.Raw.Host)
		ctx = context.WithValue(ctx, logger.Method, req.Raw.Method)
		ctx = context.WithValue(ctx, logger.URI, req.Raw.RequestURI)
		ctx = context.WithValue(ctx, logger.OperationName, req.OperationName)
		ctx = context.WithValue(ctx, logger.UserAgent, req.Raw.UserAgent())
		req.SetContext(ctx)

		tlog.Debug(ctx, "request received", &map[string]any{
			"body":       req.Body,
			"params":     req.Params,
			"host":       req.Raw.Host,
			"user_agent": req.Raw.UserAgent(),
			"uri":        req.Raw.RequestURI,
			"method":     req.Raw.Method,
			"operation":  req.OperationName,
		})

		resp, err := next(req)

		jsonBytes, merr := json.Marshal(resp.Type)

		var respBody string
		if merr != nil {
			respBody = "failed to marshal response body"
		} else {
			respBody = string(jsonBytes)
		}

		statusCode := 200

		if err != nil {
			statusCode = 500
		} else if tresp, ok := resp.Type.(interface{ GetStatusCode() int }); ok {
			statusCode = tresp.GetStatusCode()
		}

		tlog.Debug(ctx, "response sent", &map[string]any{
			"status_code": statusCode,
			"body":        respBody,
		})

		return resp, err
	}
}
