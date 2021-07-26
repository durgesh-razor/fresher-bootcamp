package Marks

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetMarks(c *gin.Context) {
	var marks []Marks
	err := GetAllMarks_(&marks)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, marks)
	}
}

//CreateUser ... Create User
func CreateMark(c *gin.Context) {
	var mark Marks
	c.BindJSON(&mark)
	err := CreateMark_(&mark)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, mark)
	}
}
