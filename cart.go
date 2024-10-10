package main

var cart = Cart{}

type CartItem struct {
	ProductID int
	Quantity  int
}
type Cart struct {
	CartItems []CartItem
}

// findProductInCart 對輸入的商品跟購物車內商品是否相同
func (c *Cart) findProductInCart(productID int) int {
	for i, item := range c.CartItems {
		if productID == item.ProductID {
			return i
		}
	}
	return -1
}

// addItem 新增商品到購物車
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

// removeItem 刪除購物車內商品
func (c *Cart) removeItem(productID int) {
	index := c.findProductInCart(productID)
	if index >= 0 {
		c.CartItems = append(c.CartItems[:index], c.CartItems[index+1:]...)
		return
	}
}
