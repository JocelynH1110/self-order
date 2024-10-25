package main

import (
	"database/sql"
)

var cart = Cart{}

type CartItem struct {
	ProductID int
	Quantity  int
}
type Cart struct {
	CartItems []CartItem
}

// findProductInCart 對輸入的商品跟購物車內商品是否相同
/*func (c *Cart) findProductInCart(productID int) int {
	for i, item := range c.CartItems {
		if productID == item.ProductID {
			return i
		}
	}
	return -1
}*/

func (c *Cart) findProductInCart(db *sql.DB, productID int) (*CartItem, error) {
	row := db.QueryRow("select id, quantity,product_id from cart_items where product_id = ?", productID)
	var i CartItem
	if err := row.Scan(&i.ProductID, &i.Quantity); err != nil {
		return nil, err
	}
	return &i, nil
}

// addItem 新增商品到購物車
/*func (c *Cart) addItem(productID, quantity int) {
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
}*/

func (c *Cart) addItem(db *sql.DB, productID, quantity int) (*CartItem, error) {
	row := db.QueryRow("insert into cart_items(id, quantity,product_id) values(?,?)")
	var i CartItem
	if err := row.Scan(&i.ProductID, &i.Quantity); err != nil {
		return nil, err
	}
	return &i, nil
}

// removeItem 刪除購物車內商品
/*func (c *Cart) removeItem(productID int) {
	index := c.findProductInCart(productID)
	if index >= 0 {
		c.CartItems = append(c.CartItems[:index], c.CartItems[index+1:]...)
	}
}*/
func (c *Cart) removeItem(db *sql.DB, productID int) (*CartItem, error) {
	row := db.QueryRow("delete from cart_items where id = ?", productID)

	var i CartItem
	if err := row.Scan(&i.ProductID, &i.Quantity); err != nil {
		return nil, err
	}
	return &i, nil
}
