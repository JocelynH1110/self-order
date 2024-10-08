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
	for _, item := range listProducts() {
		if id == item.ID {
			return &item
		}
	}
	return nil
}
