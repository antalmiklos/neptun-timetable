package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	CourseCode  string `gorm:"unique"`
	SubjectCode string
	Start       string
}

type Model interface {
	ToJSON() (string, error)
}

func (c *Course) NewCourse() *Course {
	return &Course{}
}

func (c *Course) ToJSON() (string, error) {
	return "", nil
}
