package models

import (
	"errors"
)

type Review struct {
	ID       uint   `json:"_id"`
	UserID   uint   `json:"user"`
	CourseID uint   `json:"-"`
	Rating   uint   `json:"rating"`
	Review   string `json:"review"`
	Version  string `json:"__v"`
}

func (r *Review) isValidRating() (err error) {
	if r.Rating < 0 || r.Rating > 5 {
		err = errors.New("Invalid rating")
	}
	return
}

func (r *Review) BeforeSave() (err error) {
	err = r.isValidRating()
	return
}
