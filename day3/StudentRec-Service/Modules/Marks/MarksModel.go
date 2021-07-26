package Marks

import (
	"studentRecord/Modules/Student"

	"gorm.io/gorm"
)

type Marks struct {
	gorm.Model
	Student   Student.Student `json:"student" gorm:"foreignkey:StudentID"`
	StudentID uint64
	Marks     int
	Subject   string
}

func (b *Marks) TableName() string {
	return "marks"
}
