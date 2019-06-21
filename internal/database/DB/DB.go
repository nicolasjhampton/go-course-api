package DB

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	m "github.com/nicolasjhampton/go-course-api/internal/database/models"
	"github.com/nicolasjhampton/go-course-api/internal/database/seed"
)

func Setup() (DB *gorm.DB, err error) {
	// connect
	DB, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return
	}

	// Create tables
	DB.DropTable("Courses", "Steps", "Reviews", "Users")
	DB.AutoMigrate(&m.User{}, &m.Review{}, &m.Step{}, &m.Course{})
	DB = DB.Set("gorm:auto_preload", true)

	// Seed database
	if err = seed.RunTx(DB); err != nil {
		return
	}

	return
}
