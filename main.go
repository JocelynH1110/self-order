package main

import (
	"fmt"
	"os"
)

const RESTAURANT_NAME string = "Jo's coffee shop!"

type Product struct {
	ID       int
	Name     string
	Price    int
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
func handleCommand(command string) {
	switch command {
	case "menu":
		fmt.Println("Menu：")
		lists := listProducts()
		for _, product := range lists {
			fmt.Printf("%d. %s, $%d\n", product.ID, product.Name, product.Price)
		}
	case "cart":
		fmt.Println("Your cart is currently empty.")
	case "quit":
		os.Exit(0)
	default:
		fmt.Println("Unrecognized command. Command should be one of: menu, cart, quit.")
	}
}

func main() {
	var msg string
	fmt.Printf("Welcome to %s\n", RESTAURANT_NAME)
	for {
		fmt.Print("> ")
		fmt.Scanln(&msg)
		handleCommand(msg)
	}
}
