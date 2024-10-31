package main

import (
	"database/sql"
)

// 目前 Cart 裡儲存所有項目的資料，其實只需要資料庫連結 *sql.DB，這樣就不用每次呼叫方法都要傳入 db 參數
type Cart struct {
	db *sql.DB
}

const addItemQuery = `INSERT INTO cart_items(product_id, quantity) VALUES (?,?) ON CONFLICT (product_id) DO UPDATE SET quantity = EXCLUDED.quantity+quantity, updated_at = (strftime('%s','now'));`

func (c *Cart) addItem(productID, quantity int) error {
	// please update sql to upsert
	_, err := c.db.Exec(addItemQuery, productID, quantity)
	return err
}

// 不需返回 CartItem 資料，因為不確定它是否存在過，即使存在過之後也不在
func (c *Cart) removeItem(productID int) error {
	_, err := c.db.Exec("DELETE FROM cart_items WHERE product_id = ?", productID)
	return err
}

type listProductsInCartRow struct {
	Quantity int
	Name     string
	Price    int
	Subtotal int
}

// listProductsInCart 取得購物車中的商品資料
func (c *Cart) listProductsInCart() ([]listProductsInCartRow, error) {
	rows, err := c.db.Query("SELECT ci.quantity, p.name, p.price, ci.quantity * p.price FROM cart_items ci INNER JOIN products p ON ci.product_id = p.id ORDER BY p.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []listProductsInCartRow
	for rows.Next() {
		var i listProductsInCartRow
		err := rows.Scan(&i.Quantity, &i.Name, &i.Price, &i.Subtotal)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}
