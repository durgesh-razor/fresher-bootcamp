package Routes

import (
	"studentRecord/Modules/Marks"
	"studentRecord/Modules/Student"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	Student.StudentRouter(r.Group(""))
	Marks.MarkRouter(r.Group(""))
	return r
}
