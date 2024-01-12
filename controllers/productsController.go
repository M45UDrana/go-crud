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

	c.JSON(200, gin.H{
		"product": product,
	})
}

func GetProducts(c *gin.Context) {

	var products []models.Product
	result := initializers.DB.Find(&products)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"products": products,
	})
}

func GetProductById(c *gin.Context) {

	id := c.Param("id")

	var product models.Product
	result := initializers.DB.Find(&product, id)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"product": product,
	})
}
