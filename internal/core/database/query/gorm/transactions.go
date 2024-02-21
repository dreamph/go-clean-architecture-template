package gorm

import (
	"backend/internal/core/database/query"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type dbTx struct {
	DB *gorm.DB
}

func NewDBTx(db *gorm.DB) query.DBTx {
	return &dbTx{DB: db}
}

func (t *dbTx) WithTx(ctx context.Context, fn query.TxFn) (err error) {
	tx := t.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			if e, ok := r.(error); ok {
				err = errors.Wrap(e, "rollback tx")
			}
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	err = fn(ctx, &query.AppIDB{GormDB: tx})
	if err != nil {
		return err
	}
	return nil
}
