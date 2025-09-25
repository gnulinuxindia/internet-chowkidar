package handler

import (
	"bytes"
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/spec"
)

type ApiDocsHandler interface {
	ApiDocs(ctx context.Context) (genapi.ApiDocsOK, error)
}

type apiDocsHandlerImpl struct {
}

func (h *apiDocsHandlerImpl) ApiDocs(ctx context.Context) (genapi.ApiDocsOK, error) {
	return genapi.ApiDocsOK{Data: bytes.NewReader(spec.HTML)}, nil
}
