package Marks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	MarkRouter(r.Group(""))
	return r
}

func TestGetMarksRoute(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/mark", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
