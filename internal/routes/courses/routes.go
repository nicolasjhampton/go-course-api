package courses

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	auth "github.com/nicolasjhampton/go-course-api/internal/middleware/authorization"
)

var DB *gorm.DB

func Routes(g gin.IRouter, db *gorm.DB) *gin.RouterGroup {
	DB = db
	courses := g.Group("/courses")
	{
		courses.GET("/", GetCourses)
		courses.POST("/", auth.Required(DB), CreateCourse)
	}
	course := courses.Group("/:courseid")
	{
		course.Use(FindCourse)
		course.GET("/", GetCourse)
		// course.PUT("/", UpdateCourse)
	}
	return course
}
