package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"retailer-service/Models"
	"retailer-service/Service"
)

func AddProductController(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Service.AddProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetProductbyIDController(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Service.GetProductbyID(&product, id)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetAllProductController(c *gin.Context) {
	var product []Models.Product
	err := Service.GetAllProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func UpdateProductController(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Service.UpdateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}
