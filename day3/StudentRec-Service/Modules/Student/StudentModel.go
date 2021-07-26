package Student

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name     string
	Lastname string
	DOB      string
	Address  string
}

func (b *Student) TableName() string {
	return "student"
}
