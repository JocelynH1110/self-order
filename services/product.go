package services

import (
	"database/sql"

	"github.com/jocelynh1110/self-order/types"
)

type ProductService struct {
	db *sql.DB
}

func NewProductService(db *sql.DB) ProductService {
	return ProductService{db}
}

// ListProducts 所有商品列表
func (p *ProductService) ListProducts() ([]types.Product, error) {
	rows, err := p.db.Query("select id,name,price,description from products order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []types.Product
	for rows.Next() {
		var i types.Product
		err := rows.Scan(&i.ID, &i.Name, &i.Price, &i.Description)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

// FindProductByID 比對消費者輸入的編號跟資料庫中商品是否一致
func (p *ProductService) FindProductByID(id int) (*types.Product, error) {
	// 比對輸入的商品編號是否存在
	row := p.db.QueryRow("select id, name, price, description from products where id = ?", id)
	var result types.Product
	if err := row.Scan(&result.ID, &result.Name, &result.Price, &result.Description); err != nil {
		return nil, err
	}
	return &result, nil
}
