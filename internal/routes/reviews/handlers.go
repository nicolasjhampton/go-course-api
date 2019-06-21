package reviews

import (
	"fmt"
	"github.com/gin-gonic/gin"
	m "github.com/nicolasjhampton/go-course-api/internal/database/models"
	"net/http"
)

func CreateReview(c *gin.Context) {
	var review m.Review
	var err error
	if err = c.BindJSON(&review); err == nil {
		user := c.MustGet(gin.AuthUserKey).(*m.User)
		course := c.MustGet("course").(m.Course)
		review.UserID = user.ID
		review.CourseID = course.ID
		DB.FirstOrCreate(&review, review)
		course.Reviews = append(course.Reviews, review)
		DB.Save(&course)
		redirectRoute := fmt.Sprintf("/api/v1/courses/%v", course.ID)
		c.Redirect(http.StatusCreated, redirectRoute)
	} else {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}
}
