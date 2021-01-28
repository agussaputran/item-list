package main

import (
	"Belanjaan/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/items", routes.GetItem)
	router.POST("items", routes.AddNewItem)

	router.Run()
}
