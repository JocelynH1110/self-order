package types

type Product struct {
	ID          int
	Name        string
	Price       int
	Description *string
}

type ListProductsInCartRow struct {
	Quantity int
	Name     string
	Price    int
	Subtotal int
}
