package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		if len(tokens) != 2 && len(tokens) != 3 {
			fmt.Println("USAGE: add PRODUCT_ID [QUANTITY]")
			return
		}
		productID, err := strconv.ParseInt(tokens[1], 10, 64)
		if err != nil {
			fmt.Println("USAGE: add PRODUCT_ID [QUANTITY]品項錯誤")
			return
		}
		// 如果有兩個參數，才需要解析
		var quantity int64 = 1
		if len(tokens) == 3 {
			quantity, err = strconv.ParseInt(tokens[2], 10, 64)
			if err != nil {
				fmt.Println("USAGE: add PRODUCT_ID [QUANTITY]數量錯誤")
				return
			}
		}
		cart.addItem(int(productID), int(quantity))
	case "cart":
		fmt.Println("Your cart is currently empty.")
		fmt.Printf("%#v\n", cart.CartItems)
	//	fmt.Print("add %v ")
	case "quit":
		os.Exit(0)
	default:
		fmt.Println("Unrecognized command. Command should be one of: menu, cart, add, quit.")
	}
}

func main() {
	fmt.Printf("Welcome to %s\n", RESTAURANT_NAME)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		next := scanner.Scan()
		if !next {
			fmt.Println("\nBye~")
			return
		}
		line := scanner.Text()
		handleCommand(line)
	}
}
