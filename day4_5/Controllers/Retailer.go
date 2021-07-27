package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"retailer-service/Models"
	"retailer-service/Service"
)

func TransactionController(c *gin.Context) {
	var orders []Models.Orders
	c.BindJSON(&orders)
	err := Service.Transaction(&orders)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, orders)
	}
}
