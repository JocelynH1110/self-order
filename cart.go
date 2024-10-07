package main

var cart = Cart{}

type CartItem struct {
	ProductID int
	Quantity  int
}
type Cart struct {
	CartItems []CartItem
}

func (c *Cart) findProductInCart(productID int) int {
	for i, item := range c.CartItems {
		if productID == item.ProductID {
			return i
		}
	}
	return -1
}

func (c *Cart) addItem(productID, quantity int) {
	// check if the product is already in cart
	// if not, add a new cartitem
	// if yes, add quantity
	index := c.findProductInCart(productID)
	if index >= 0 {
		c.CartItems[index].Quantity += quantity
		return
	}
	item := CartItem{
		ProductID: productID,
		Quantity:  quantity,
	}
	c.CartItems = append(c.CartItems, item)
}
