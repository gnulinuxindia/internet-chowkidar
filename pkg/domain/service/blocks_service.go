package service

import (
	"context"
	"time"

	genapi "github.com/gnulinuxindia/internet-chowkidar/api/gen"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/dto"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/repository"
	"github.com/go-errors/errors"
)

type BlocksService interface {
	GetAllBlocks(ctx context.Context) ([]genapi.Block, error)
	CreateBlock(ctx context.Context, req *genapi.BlockInput) (*genapi.Block, error)
}

type blocksServiceImpl struct {
	blocksRepo repository.BlocksRepository
}

func (b *blocksServiceImpl) GetAllBlocks(ctx context.Context) ([]genapi.Block, error) {
	panic("not implemented")
}

func (b *blocksServiceImpl) CreateBlock(ctx context.Context, req *genapi.BlockInput) (*genapi.Block, error) {
	blockDto := &dto.BlockDto{	
		IspID: req.IspID,
		SiteID: req.SiteID,
		IsBlocked: req.IsBlocked,
		LastReportedAt: time.Now(),
	}

	eb, err := b.blocksRepo.CreateBlock(ctx, blockDto)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	block := &genapi.Block{
		ID: genapi.NewOptInt(eb.ID),
		IspID: genapi.NewOptInt(eb.IspID),
		SiteID: genapi.NewOptInt(eb.SiteID),
		LastReportedAt: genapi.NewOptDateTime(eb.LastReportedAt),
		BlockReports: genapi.NewOptInt(eb.BlockReports),
		UnblockReports: genapi.NewOptInt(eb.UnblockReports),
		CreatedAt: genapi.NewOptDateTime(eb.CreatedAt),
		UpdatedAt: genapi.NewOptDateTime(eb.UpdatedAt),
	}

	return block, nil
}
