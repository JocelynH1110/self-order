package main

import (
	"database/sql"
)

var cart = Cart{}

type CartItem struct {
	ProductID int
	Quantity  int
}

// 目前 Cart 裡儲存所有項目的資料，其實只需要資料庫連結 *sql.DB，這樣就不用每次呼叫方法都要傳入 db 參數
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

// 用不到因為檢查商品是否存在於購物車中的方法，改寫在資料庫中
func (c *Cart) findProductInCart(db *sql.DB, productID int) (*CartItem, error) {
	row := db.QueryRow("select id, quantity,product_id from cart_items where product_id = ?", productID)
	var i CartItem
	if err := row.Scan(&i.ProductID, &i.Quantity); err != nil {
		return nil, err
	}
	return &i, nil
}

// addItem 新增商品到購物車
// 原本用在檢查商品是否以在購物車中，現改於在 sql 中使用 upsert ，也就是說 insert ... on conflict do update set ...
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

// 不需返回 CartItem 資料，因為不確定它是否存在過，即使存在過之後也不在
func (c *Cart) removeItem(db *sql.DB, productID int) error {
	row := db.QueryRow("delete from cart_items where product_id = ?", productID)

	var i CartItem
	if err := row.Scan(&i.ProductID, &i.Quantity); err != nil {
		return nil, err
	}
	return &i, nil
}

// listProductsInCart 取得購物車中的商品資料
// select ci.quantity,p.name,p.price,ci.quantity*p.price FROM cart_items ci INNER JOIN products p ON CI.product_id =p.id ORDER BY p.id ;
