package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Step struct {
	ID          uint   `json:"_id"`
	CourseID    uint   `json:"-"`
	Number      uint   `json:"stepNumber"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (s *Step) isRequired() (err error) {
	if len(s.Title) <= 0 || len(s.Description) <= 0 {
		err = errors.New("Required Step field missing")
	}
	return
}

func (s *Step) BeforeSave() (err error) {
	err = s.isRequired()
	return
}

func (s *Step) AfterCreate(tx *gorm.DB) (err error) {
	var course Course
	tx.Table("courses").First(&course, s.CourseID)
	tx.Model(s).Update("Number", len(course.Steps))
	return
}
