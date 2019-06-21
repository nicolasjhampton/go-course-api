package seed

import (
	"github.com/jinzhu/gorm"
	m "github.com/nicolasjhampton/go-course-api/internal/database/models"
)

var users = []m.User{
	m.User{
		ID:       3,
		Name:     "Joe Smith",
		Email:    "joe@smith.com",
		Password: "password",
	},
	m.User{
		ID:       2,
		Name:     "Sam Jones",
		Email:    "sam@jones.com",
		Password: "password",
	},
	m.User{
		ID:       1,
		Name:     "Sam Smith",
		Email:    "sam@smith.com",
		Password: "password",
	},
}

func userseeds(tx *gorm.DB) (err error) {
	for _, user := range users {
		if err = tx.Create(&user).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	return
}
