package Marks

import "github.com/gin-gonic/gin"

func MarkRouter(router *gin.RouterGroup) {

	router.GET("mark", GetMarks)
	router.POST("mark", CreateMark)
	// router.GET("mark/:id", CreateStudent)
	// router.PUT("mark/:id", UpdateStudent)
	// router.DELETE("mark/:id", DeleteStudent)
}
