package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Record struct {
	ID    int
	Model string
	// Add other fields as needed
}

func searchByTitle(db *sql.DB, searchStr string) ([]Record, error) {
	query := "SELECT id, model FROM buybacks WHERE model LIKE ?"
	searchPattern := "%" + searchStr + "%"

	rows, err := db.Query(query, searchPattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []Record

	for rows.Next() {
		var record Record
		err := rows.Scan(&record.ID, &record.Model)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func main() {
	// Replace with your MySQL connection details
	dsn := "root:password@tcp(localhost:3306)/marketdata_test"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	searchStr := "Phone"
	records, err := searchByTitle(db, searchStr)
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		fmt.Printf("ID:")
		fmt.Printf("ID: %d, Model: %s\n", record.ID, record.Model)
	}
}
