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
	url, present := os.LookupEnv("DATABASE_URL")
	if !present {
		url = "postgresql://localhost/course-api?sslmode=disable"
	}
	DB, err = gorm.Open("postgres", url)
	if err != nil {
		return
	}

	// Create tables
	DB.Exec("DROP TABLE courses;")
	DB.Exec("DROP TABLE steps;")
	DB.Exec("DROP TABLE reviews;")
	DB.Exec("DROP TABLE users;")
	// DB.DropTable("Courses", "Steps", "Reviews", "Users")
	DB.AutoMigrate(&m.User{}, &m.Review{}, &m.Step{}, &m.Course{})
	DB = DB.Set("gorm:auto_preload", true)

	// Seed database
	if err = seed.RunTx(DB); err != nil {
		return
	}

	return
}
