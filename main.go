package main

import (
	"github.com/M45UDrana/go-crud/controllers"
	"github.com/M45UDrana/go-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.CreateProduct)

	r.Run()
}
