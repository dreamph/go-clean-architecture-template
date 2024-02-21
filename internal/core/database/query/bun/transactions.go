package bun

import (
	"backend/internal/core/database/query"
	"context"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

type dbTx struct {
	DB *bun.DB
}

func NewDBTx(db *bun.DB) query.DBTx {
	return &dbTx{DB: db}
}

func (t *dbTx) WithTx(ctx context.Context, fn query.TxFn) (err error) {
	tx, err := t.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			if e, ok := r.(error); ok {
				err = errors.Wrap(e, "rollback tx")
			}
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(ctx, &query.AppIDB{BunDB: tx})
	if err != nil {
		return err
	}
	return nil
}
