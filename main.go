package main

import (
	"fmt"
	"os"
	"strings"
)

const RESTAURANT_NAME string = "Jo's coffee shop!"

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
func add(id string, quantity int) {

}
func handleCommand(command string) {
	tokens := strings.Split(command, " ")
	switch tokens[0] {
	case "menu":
		fmt.Println("Menu：")
		lists := listProducts()
		for _, product := range lists {
			fmt.Printf("%d. %s, $%d\n", product.ID, product.Name, product.Price)
		}
	case "add":
	case "cart":
		fmt.Println("Your cart is currently empty.")
	//	fmt.Print("add %v ")
	case "quit":
		os.Exit(0)
	default:
		fmt.Println("Unrecognized command. Command should be one of: menu, cart, add, quit.")
	}
}

func main() {
	var msg string
	fmt.Printf("Welcome to %s\n", RESTAURANT_NAME)
	for {
		fmt.Print("> ")
		fmt.Scanln(&msg)
		handleCommand(msg)
		if msg == "cart" {
			a := CartItem{}
			fmt.Print("add %v", a.ProductID)
		}

	}
}
