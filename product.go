package main

type Product struct {
	ID    int
	Name  string
	Price int
}

func listProducts() []Product {
	return []Product{
		{
			ID:    1,
			Name:  "Latte",
			Price: 150,
		},
		{
			ID:    2,
			Name:  "Filter coffee",
			Price: 100,
		},
	}
}
func findProductByID(id int) *Product {
	// 比對輸入的商品編號是否存在
	for _, item := range listProducts() {
		if id == item.ID {
			return &item
		}
	}
	return nil
}
