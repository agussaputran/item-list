package main

import (
	"Belanjaan/route"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	var itemList route.ItemList
	fmt.Println(itemList.ListOfItem)

	router.POST("/items", itemList.AddNewItem)
	router.GET("/items", itemList.GetItem)

	router.Run()
}
