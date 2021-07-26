package Student

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetStudents(c *gin.Context) {
	var students []Student
	err := GetAllStudents_(&students)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, students)
	}
}

//CreateUser ... Create User
func CreateStudent(c *gin.Context) {
	var student Student
	c.BindJSON(&student)
	err := CreateStudent_(&student)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//GetUserByID ... Get the user by id
func GetStudentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var student Student
	err := GetStudentByID_(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//UpdateUser ... Update the user information
func UpdateStudent(c *gin.Context) {
	var student Student
	id := c.Params.ByName("id")
	err := GetStudentByID_(&student, id)
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}
	c.BindJSON(&student)
	err = UpdateStudent_(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//DeleteUser ... Delete the user
func DeleteStudent(c *gin.Context) {
	var student Student
	id := c.Params.ByName("id")
	err := DeleteStudent_(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
