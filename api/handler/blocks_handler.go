package handler

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
)

type BlocksHandler interface {
	ListBlocks(ctx context.Context) ([]genapi.Block, error)
	CreateBlock(ctx context.Context, req *genapi.BlockInput) (*genapi.Block, error)
}

type blocksHandlerImpl struct {
}

func (b *blocksHandlerImpl) ListBlocks(ctx context.Context) ([]genapi.Block, error) {
	panic("not implemented")
}

func (b *blocksHandlerImpl) CreateBlock(ctx context.Context, req *genapi.BlockInput) (*genapi.Block, error) {
	panic("not implemented")
}
