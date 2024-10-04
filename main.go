package main

import "fmt"

var restaurant string = "Jo's coffee shop!"

type List struct {
	productID   int
	productName string
	price       int
}

func getList() []List {
	return []List{
		{
			productID:   1,
			productName: "Latte",
			price:       150,
		},
		{
			productID:   2,
			productName: "Filter coffee",
			price:       100,
		},
	}
}
func handleCommand(command string) {
	switch command {
	case "menu":
		fmt.Println("Menu：")
		lists := getList()
		for _, product := range lists {
			fmt.Printf("%v. %v, $%v\n", product.productID, product.productName, product.price)
		}
	case "cart":
		fmt.Println("Your cart is currently empty.")
	default:
		fmt.Println("Command should be one of: menu, cart, quit.")
	}
}
func main() {
	var msg string
	fmt.Printf("Welcome to %s\n", restaurant)
	fmt.Scanln(&msg)
	for {
		handleCommand(msg)
		fmt.Scanln(&msg)
	}
}
