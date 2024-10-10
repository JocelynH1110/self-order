package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"strings"
)

const RESTAURANT_NAME string = "Jo's coffee shop!"

func parseRemoveCommand(tokens []string) (ok bool, productID int) {
	if len(tokens) != 2 {
		return
	}
	id, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil || id < 1 {
		return
	}
	return true, int(id)
}
func parseAddCommand(tokens []string) (ok bool, productID int, quantity int) {
	if len(tokens) != 2 && len(tokens) != 3 {
		return
	}
	id, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil || id < 1 {
		return
	}
	if len(tokens) == 2 {
		return true, int(id), 1
	}
	qty, err := strconv.ParseInt(tokens[2], 10, 64)
	if err != nil || qty < 1 {
		return
	}
	return true, int(id), int(qty)
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
	case "remove":
		if ok, productID := parseRemoveCommand(tokens); ok {
			cart.removeItem(productID)
			return
		}
		fmt.Println("USAGE: add PRODUCT_ID [QUANTITY]")
	case "add":
		if ok, productID, quantity := parseAddCommand(tokens); ok {
			cart.addItem(productID, quantity)
			return
		}
		fmt.Println("USAGE: add PRODUCT_ID [QUANTITY]")

	case "cart":
		if cart.CartItems == nil {
			fmt.Println("Your cart is currently empty.")
			return
		}
		fmt.Printf("You have %v item(s) in the shopping cart.\n", len(cart.CartItems))

		// 顯示已輸入並存在的商品列表
		for index, item := range cart.CartItems {
			product := findProductByID(item.ProductID)
			if product == nil {
				continue
			}
			fmt.Printf("%d. %s, %d x $%d = $%d\n", index+1, product.Name, item.Quantity, product.Price, item.Quantity*product.Price)
		}

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
