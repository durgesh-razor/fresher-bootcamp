package Student

import (
	"fmt"
	"studentRecord/Config"
)

func GetAllStudents_(students *[]Student) (err error) {
	if err = Config.DB.Find(students).Error; err != nil {
		return err
	}
	return nil
}

func CreateStudent_(Student *Student) (err error) {
	if err = Config.DB.Create(Student).Error; err != nil {
		return err
	}
	return nil
}

func GetStudentByID_(Student *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(Student).Error; err != nil {
		return err
	}
	return nil
}

func UpdateStudent_(Student *Student, id string) (err error) {
	fmt.Println(Student)
	Config.DB.Save(Student)
	return nil
}

func DeleteStudent_(Student *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(Student)
	return nil
}
