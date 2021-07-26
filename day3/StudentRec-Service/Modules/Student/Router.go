package Student

import "github.com/gin-gonic/gin"

func UserRouter(router *gin.RouterGroup) {

	router.GET("student", GetStudents)
	router.POST("student", CreateStudent)
	router.GET("student/:id", GetStudentByID)
	router.PUT("student/:id", UpdateStudent)
	router.DELETE("student/:id", DeleteStudent)
}
