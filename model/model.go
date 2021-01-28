package model

// ItemModel struct
type ItemModel struct {
	Name  string  `json:"name"`
	Qty   int     `json:"qty"`
	Price float32 `json:"price"`
}

// ItemList struct
type ItemList struct {
	ListOfItem []ItemModel
}
