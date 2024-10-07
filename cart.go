package main

import "fmt"

type CartItem struct {
	ProductID int
	Quantity  int
}
type Cart struct {
	CartItems []CartItem
}

func (c *CartItem) setCartItem(productID, quantity int) {
	c.ProductID = productID
	c.Quantity = quantity
}
func (c CartItem) getCartItem() string {
	return fmt.Sprintf("%d,%d", c.ProductID, c.Quantity)
}
