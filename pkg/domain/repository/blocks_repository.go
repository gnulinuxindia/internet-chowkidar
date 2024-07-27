package repository

import (
	"context"
	"log/slog"

	"github.com/gnulinuxindia/internet-chowkidar/ent"
	"github.com/gnulinuxindia/internet-chowkidar/ent/blocks"
	"github.com/gnulinuxindia/internet-chowkidar/pkg/domain/dto"
)

type BlocksRepository interface {
	GetAllBlocks(ctx context.Context) ([]*ent.Blocks, error)
	CreateBlock(ctx context.Context, req *dto.BlockDto) (*ent.Blocks, error)
}

type blocksRepositoryImpl struct {
	db *ent.Client
}

func (b *blocksRepositoryImpl) GetAllBlocks(ctx context.Context) ([]*ent.Blocks, error) {
	return b.db.Blocks.Query().All(ctx)
}

func (b *blocksRepositoryImpl) CreateBlock(ctx context.Context, req *dto.BlockDto) (*ent.Blocks, error) {
	tx, err := b.db.Tx(ctx)
	if err != nil {
		slog.Error("failed to create transaction", "err", err)
		return nil, err
	}

	var blk *ent.Blocks

	// check if block already exists
	eBlk, err := tx.Blocks.Query().
		Where(
			blocks.IspIDEQ(req.IspID),
			blocks.SiteIDEQ(req.SiteID),
		).
		First(ctx)
	if eBlk != nil {
		slog.Info("block already exists", "block", eBlk)

		uQuery := tx.Blocks.UpdateOne(eBlk).
			SetLastReportedAt(req.LastReportedAt)

		if req.IsBlocked {
			uQuery.SetBlockReports(eBlk.BlockReports + 1)
		} else {
			uQuery.SetUnblockReports(eBlk.UnblockReports + 1)
		}

		blk, err = uQuery.Save(ctx)
		if err != nil {
			slog.Error("failed to update block", "err", err)
			return nil, rollback(tx, err)
		}

		slog.Info("block updated", "block", blk)
	} else {
		// create new block
		iQuery := tx.Blocks.Create().
			SetIspID(req.IspID).
			SetSiteID(req.SiteID).
			SetLastReportedAt(req.LastReportedAt)

		if req.IsBlocked {
			iQuery.SetBlockReports(1)
		} else {
			iQuery.SetUnblockReports(1)
		}

		blk, err = iQuery.Save(ctx)
		if err != nil {
			slog.Error("failed to create block", "err", err)
			return nil, rollback(tx, err)
		}

		slog.Info("block created", "block", blk)
	}

	return blk, tx.Commit()
}
