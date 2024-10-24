package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	ID          int
	Name        string
	Price       int
	Description *string
}

func listProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id,name,price,description from products order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Product
	for rows.Next() {
		var i Product
		err := rows.Scan(&i.ID, &i.Name, &i.Price, &i.Description)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

func main() {
	db, err := sql.Open("sqlite3", "db/dev.db")
	if err != nil {
		log.Fatalf("ERROR initializing database: %s", err)
	}
	defer db.Close()

	row := db.QueryRow("select 2+2")
	var result int
	if err := row.Scan(&result); err != nil {
		log.Fatal(err)
	}
	products, err := listProducts(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", products)
}
