package courses

import (
	"fmt"
	"github.com/gin-gonic/gin"
	m "github.com/nicolasjhampton/go-course-api/internal/database/models"
	"net/http"
)

func GetCourses(c *gin.Context) {
	var results []struct {
		ID    uint   `json:"_id"`
		Title string `json:"title"`
	}
	DB.Table("courses").Select("ID, Title").Scan(&results)
	c.JSON(http.StatusOK, results)
}

func FindCourse(c *gin.Context) {
	var course m.Course
	id := c.Param("courseid")
	DB.First(&course, id)
	c.Set("course", course)
}

func GetCourse(c *gin.Context) {
	course := c.MustGet("course").(m.Course)
	c.JSON(http.StatusOK, course)
}

func CreateCourse(c *gin.Context) {
	var course m.Course
	var err error
	var count int
	var status int
	if err = c.BindJSON(&course); err == nil {
		user := c.MustGet(gin.AuthUserKey).(*m.User)
		courses := DB.Table("courses").Where("Title = ?", course.Title)
		if courses.Count(&count); count == 0 {
			status = http.StatusCreated
		} else {
			status = http.StatusFound
		}
		courses.Assign(m.Course{UserID: user.ID}).FirstOrCreate(&course)
		c.Redirect(status, fmt.Sprintf("/api/v1/courses/%v", course.ID))
	} else {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
	}
}
