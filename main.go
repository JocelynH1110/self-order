package main

import "fmt"

type List struct {
	productID   int
	productName string
	price       int
}

func getList() []List {
	l1 := List{
		productID:   1,
		productName: "Latte",
		price:       150,
	}
	l2 := List{
		productID:   2,
		productName: "Filter coffee",
		price:       100,
	}
	return []List{l1, l2}
}
func main() {
	lists := getList()
	for i := 0; i < len(lists); i++ {
		fmt.Println(lists[i])
	}
}
