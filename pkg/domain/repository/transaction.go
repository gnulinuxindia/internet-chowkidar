package repository

import (
	"context"
	"fmt"

	"github.com/gnulinuxindia/internet-chowkidar/ent"
)

const TxKey = "tx"

type TxHandler struct {
	client *ent.Client
}

func NewTxHandler(client *ent.Client) *TxHandler {
	return &TxHandler{client: client}
}

func (t *TxHandler) Begin(ctx *context.Context) error {
	tx, err := t.client.Tx(*ctx)
	if err != nil {
		return err
	}

	(*ctx) = context.WithValue(*ctx, TxKey, tx)
	return nil
}

func (t *TxHandler) GetTxFromCtx(ctx context.Context) (*ent.Tx, error) {
	tx, ok := ctx.Value(TxKey).(*ent.Tx)
	if !ok {
		return nil, fmt.Errorf("no transaction found in context")
	}
	return tx, nil
}

func (t *TxHandler) CommitOrRollback(ctx context.Context) error {
	tx, err := t.GetTxFromCtx(ctx)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return rollback(tx, err)
	}
	return nil
}

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
