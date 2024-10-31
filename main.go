package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

const RESTAURANT_NAME string = "Jo's coffee shop!"

/*func parseRemoveCommand(tokens []string) (ok bool, productID int) {
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
*/

func menuHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lists, err := listProducts(db)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, `<title>%s</title><h1>%s</h1><h2>Menu</h2><ol>`, RESTAURANT_NAME, RESTAURANT_NAME)
		for _, product := range lists {
			fmt.Fprintf(w, "<li>%s, $%d</li>\n", product.Name, product.Price)
		}

	})
}

/*
	func handleCommand(db *sql.DB, command string) {
		cart := Cart{db} // 現在 cart 依賴資料庫連接，不能存在全域變數內，故可在 handleCommand 內初始化
		tokens := strings.Split(command, " ")
		switch tokens[0] {
		case "menu":
			fmt.Println("Menu：")
			lists, err := listProducts(db)
			if err != nil {
				log.Fatal(err)
			}
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
			items, err := cart.listProductsInCart()
			if err != nil {
				log.Fatal(err)
			}
			if items == nil {
				fmt.Println("Your cart is currently empty.")
				return
			}
			fmt.Printf("You have %v item(s) in the shopping cart.\n", len(items))
			for index, item := range items {
				fmt.Printf("%d. %s, %d x $%d = $%d\n", index+1, item.Name, item.Quantity, item.Price, item.Subtotal)
			}

		case "quit":
			os.Exit(0)

		default:
			fmt.Println("Unrecognized command. Command should be one of: menu, cart, add, quit.")
		}
	}
*/
func main() {
	db, err := sql.Open("sqlite3", "db/dev.db")
	if err != nil {
		log.Fatalf("ERROR initializing database: %s", err)
	}
	defer db.Close()

	http.Handle("/", menuHandler(db))
	log.Fatal(http.ListenAndServe(":3000", nil))

	/*fmt.Printf("Welcome to %s\n", RESTAURANT_NAME)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		next := scanner.Scan()
		if !next {
			fmt.Println("\nBye~")
			return
		}
		line := scanner.Text()
		handleCommand(db, line)
	}*/
}
