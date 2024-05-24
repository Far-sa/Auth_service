package models

type CartItem struct {
	ProductID string
	Quantity  int32
}

type Cart struct {
	UserID string
	Items  []CartItem
}
