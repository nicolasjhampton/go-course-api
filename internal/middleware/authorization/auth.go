package authorization

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "github.com/nicolasjhampton/go-course-api/internal/database/models"
	"net/http"
	"regexp"
	"strings"
)

type auth struct {
	Email string
	Pass  string
}

var admins []string = []string{
	"sam@smith.com",
}

var re *regexp.Regexp = regexp.MustCompile(`Basic\s([^\s]+)`)

func Admin(DB *gorm.DB) gin.HandlerFunc {
	return authorize(DB, true)
}

func Required(DB *gorm.DB) gin.HandlerFunc {
	return authorize(DB, false)
}

func authorize(DB *gorm.DB, needsAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *m.User
		var auth *auth
		var err error

		if auth, err = parseAuthHeader(c); err == nil {
			if user, err = getAuthorizedUser(auth, DB); err == nil {
				err = checkAdmin(user, needsAdmin)
			}
		}

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.Set(gin.AuthUserKey, user)
		}
	}
}

func parseAuthHeader(c *gin.Context) (*auth, error) {
	reAuth := re.FindStringSubmatch(c.GetHeader("Authorization"))
	if len(reAuth) < 2 {
		return nil, errors.New("No vaild authorization header found")
	}

	authHead, err := base64.StdEncoding.DecodeString(reAuth[1])
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Incorrect base64 encoding")
	}

	authArr := strings.Split(string(authHead), ":")
	if len(authArr) < 2 {
		return nil, errors.New("Authorization header malformed")
	}

	return &auth{Email: authArr[0], Pass: authArr[1]}, nil

}

func getAuthorizedUser(auth *auth, DB *gorm.DB) (*m.User, error) {
	var user m.User
	var err error
	DB.Where(&m.User{Email: auth.Email}).First(&user)
	if auth.Email != user.Email || auth.Pass != user.Password {
		err = errors.New("authentication failure")
	}
	return &user, err
}

func checkAdmin(user *m.User, needsAdmin bool) error {
	if !needsAdmin {
		return nil
	}
	var authorized = false
	for _, admin := range admins {
		if admin == user.Email {
			authorized = true
			break
		}
	}
	if !authorized {
		return errors.New("Admin only route")
	}
	return nil
}
