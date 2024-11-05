package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jocelynh1110/self-order/services"
	"github.com/jocelynh1110/self-order/types"
	_ "github.com/mattn/go-sqlite3"
)

const RESTAURANT_NAME string = "Jo's coffee shop!"

type MenuProps struct {
	Title    string
	Products []types.Product
}

var menuTemplate = template.Must(template.ParseFiles("./templates/index.html.tmpl"))

func menuHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ps := services.NewProductService(db)
		lists, err := ps.ListProducts()
		if err != nil {
			log.Fatal(err)
		}
		menuTemplate.Execute(w, MenuProps{
			Title:    RESTAURANT_NAME,
			Products: lists,
		})
	}
}

func cartHandler(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cs := services.NewCartService(db)
		items, err := cs.ListProductsInCart()
		if err != nil {
			log.Fatal(err)
		}
		if items == nil {
			fmt.Fprintf(w, "Your cart is currently empty.")
			return
		}
		fmt.Fprintf(w, "You have %v item(s) in the shopping cart.\n", len(items))
		for index, item := range items {
			fmt.Fprintf(w, "%d. %s, %d x $%d = $%d\n", index+1, item.Name, item.Quantity, item.Price, item.Subtotal)
		}
	})
}

func main() {
	db, err := sql.Open("sqlite3", "db/dev.db")
	if err != nil {
		log.Fatalf("ERROR initializing database: %s", err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger) // 在終端機會出現日誌資料等
	r.Get("/", menuHandler(db))
	r.Get("/items", cartHandler(db))
	log.Fatal(http.ListenAndServe(":3000", r))
}
