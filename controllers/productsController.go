package controllers

import (
	"github.com/M45UDrana/go-crud/initializers"
	"github.com/M45UDrana/go-crud/models"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {

	var body struct {
		Code  string
		Price uint
	}

	c.Bind(&body)

	product := models.Product{Code: body.Code, Price: body.Price}
	result := initializers.DB.Create(&product)
	if result.Error != nil {
		c.Status(400)
		return
	}
	//result.RowsAffected

	c.JSON(200, gin.H{
		"product": product,
	})
}
