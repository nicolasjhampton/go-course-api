package users

import (
	"github.com/gin-gonic/gin"
	m "github.com/nicolasjhampton/go-course-api/internal/database/models"
	"net/http"
)

func GetUser(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(*m.User)
	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	var users []m.User
	DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user m.User
	var err error
	if err = c.BindJSON(&user); err == nil {
		DB.FirstOrCreate(&user, user)
		c.JSON(http.StatusFound, user)
	} else {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}
}
