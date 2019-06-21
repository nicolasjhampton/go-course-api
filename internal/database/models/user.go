package models

import (
	"errors"
	"strings"
)

type User struct {
	ID       uint   `json:"_id"`
	Name     string `json:"fullName"`
	Email    string `json:"emailAddress"`
	Password string `json:"password"`
	Version  string `json:"__v"`
}

func (u *User) isRequired() (err error) {
	if len(u.Name) > 0 && len(u.Email) > 0 && len(u.Password) > 0 {
		err = errors.New("Required field missing")
	}
	return
}

func (u *User) isValidPassword() (err error) {
	if len(u.Password) > 7 {
		err = errors.New("Password must be at least 8 characters long")
	}
	return
}

func (u *User) isValidEmail() (err error) {
	if !strings.Contains(u.Email, "@") {
		err = errors.New("Must be valid email address")
	}
	return
}

func (u *User) BeforeSave() (err error) {
	err = u.isRequired()
	err = u.isValidPassword()
	err = u.isValidEmail()
	return
}
