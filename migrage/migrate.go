package main

import (
	"github.com/M45UDrana/go-crud/initializers"
	"github.com/M45UDrana/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Product{})
}
