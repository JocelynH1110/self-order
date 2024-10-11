package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func initDatabase(dbPath string) error {
	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	db = conn

	// create products table
	_, err = conn.Exec(
		`CREATE TABLE IF NOT EXISTS products (
				id INTEGER PRIMARY KEY AUTOINCREMENT, 
				name VARCHAR(20) NOT NULL, 
				price INTEGER NOT NULL,
			check(price>=0)
		);`,
	)
	return err
}

// insert a item into the database
func addListProduct(p *Product) error {
	_, err := db.Exec(
		`INSERT INTO products (ID, Name, Price) VALUES (?,?,?);`, p.ID, p.Name, p.Price,
	)
	return err
}
func main() {
	err := initDatabase(":memory:")
	if err != nil {
		log.Fatalf("ERROR initializing database: %s", err)
	}
	defer db.Close()

	row := db.QueryRow("select 2+2")
	var result int
	if err := row.Scan(&result); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

}
