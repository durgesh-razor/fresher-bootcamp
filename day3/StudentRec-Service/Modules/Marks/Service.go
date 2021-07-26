package Marks

import (
	"fmt"
	"studentRecord/Config"
)

func GetAllMarks_(marks *[]Marks) (err error) {
	if err = Config.DB.Find(marks).Error; err != nil {
		return err
	}
	return nil
}

func CreateMark_(mark *Marks) (err error) {
	fmt.Println(mark)
	if err = Config.DB.Create(mark).Error; err != nil {
		// panic(err)
		return err
	}
	return nil
}
