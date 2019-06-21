package models

import (
	"errors"
)

type Course struct {
	ID          uint     `json:"_id"`
	UserID      uint     `json:"-"`
	User        User     `json:"user"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Time        string   `json:"estimatedTime"`
	Materials   string   `json:"materialsNeeded"`
	Reviews     []Review `json:"reviews"`
	Steps       []Step   `json:"steps"`
	Version     string   `json:"__v"`
}

func (c *Course) isRequired() (err error) {
	if len(c.Title) <= 0 || len(c.Description) <= 0 || len(c.Time) <= 0 || len(c.Materials) <= 0 {
		err = errors.New("Required Course field missing")
	}
	return
}

func (c *Course) BeforeSave() (err error) {
	err = c.isRequired()
	return
}
