package seed

import (
	"github.com/jinzhu/gorm"
)

type SeedFunc func(*gorm.DB) error

var seeds = []SeedFunc{
	userseeds,
	courseseeds,
}

func RunTx(DB *gorm.DB) (err error) {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Error; err != nil {
		return
	}

	for _, s := range seeds {
		if err = s(tx); err != nil {
			return
		}
	}

	err = tx.Commit().Error
	return
}
