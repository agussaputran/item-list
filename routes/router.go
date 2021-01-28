package routes

import (
	"Belanjaan/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

// AddNewItem Example
func AddNewItem(c *gin.Context) {
	var item model.ItemModel
	var itemList *model.ItemList
	err := c.Bind(&item)
	if err != nil {
		c.JSON(400, gin.H{
			"err_code": 400,
			"msg":      err,
		})
	}

	c.JSON(200, gin.H{
		"item":     item.Name,
		"quantity": item.Qty,
		"price":    item.Price,
	})
	itemList.ListOfItem = append(itemList.ListOfItem, item)
}

//GetItem example
func GetItem(c *gin.Context) {
	var itemList *model.ItemList
	err := c.Bind(&itemList.ListOfItem)
	if err != nil {
		fmt.Println("ERROR")
	}

	c.JSON(200, gin.H{
		"data": itemList.ListOfItem,
	})
}
