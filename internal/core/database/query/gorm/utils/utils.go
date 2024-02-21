package utils

import (
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

func DbError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	return err
}
