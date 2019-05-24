package db

import (
	"github.com/jinzhu/gorm"
)

func FinishTx(tx *gorm.DB, err error) error {
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
