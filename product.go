package main

import (
	"database/sql"
)

type Product struct {
	ID          int
	Name        string
	Price       int
	Description *string
}

// listProducts 所有商品列表
func listProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id,name,price,description from products order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Product
	for rows.Next() {
		var i Product
		err := rows.Scan(&i.ID, &i.Name, &i.Price, &i.Description)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

// findProductByID 比對消費者輸入的編號跟資料庫中商品是否一致
func findProductByID(db *sql.DB, id int) (*Product, error) {
	// 比對輸入的商品編號是否存在
	row := db.QueryRow("select id, name, price, description from products where id = ?", id)
	var result Product
	if err := row.Scan(&result.ID, &result.Name, &result.Price, &result.Description); err != nil {
		return nil, err
	}
	return &result, nil
}
