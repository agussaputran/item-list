package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// ItemModel struct
type ItemModel struct {
	Name  string `json:"name"`
	Qty   string `json:"qty"`
	Price string `json:"price"`
}

// ItemList struct
type ItemList struct {
	ListOfItem []ItemModel
}

//AddNewItem example
func (list *ItemList) AddNewItem(c *gin.Context) {
	var items ItemModel
	itemName := c.PostForm("name")
	itemQty := c.PostForm("qty")
	itemPrice := c.PostForm("price")

	items.Name = itemName
	items.Qty = itemQty
	items.Price = itemPrice

	list.ListOfItem = append(list.ListOfItem, items)

	c.JSON(200, gin.H{
		"item":  itemName,
		"qty":   itemQty,
		"price": itemPrice,
	})

}

//GetItem example
func (list *ItemList) GetItem(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": list.ListOfItem,
	})
}

//UpdateItem example
func (list *ItemList) UpdateItem(c *gin.Context) {
	var nameInput = c.Param("item")

	// var updatedName = c.PostForm("name")
	var updatedQty = c.PostForm("qty")
	var updatedPrice = c.PostForm("price")

	for _, v := range list.ListOfItem {
		fmt.Println(v.Name)
		if nameInput == v.Name {
			// v.Name = updatedName
			v.Qty = updatedQty
			v.Price = updatedPrice
		}
	}
}
