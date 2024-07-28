package handler

import (
	"context"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/service"
	"github.com/go-errors/errors"
)

type BlocksHandler interface {
	ListBlocks(ctx context.Context) ([]genapi.Block, error)
	CreateBlock(ctx context.Context, req *genapi.BlockInput) (*genapi.Block, error)
}

type blocksHandlerImpl struct {
	blocksService service.BlocksService
	sitesService  service.SitesService
}

func (b *blocksHandlerImpl) ListBlocks(ctx context.Context) ([]genapi.Block, error) {
	blocks, err := b.blocksService.GetAllBlocks(ctx)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return blocks, nil
}

func (b *blocksHandlerImpl) CreateBlock(ctx context.Context, req *genapi.BlockInput) (*genapi.Block, error) {
	block, err := b.blocksService.CreateBlock(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return block, nil
}
