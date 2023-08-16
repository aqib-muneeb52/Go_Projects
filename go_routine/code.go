package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type Record struct {
	ID    int
	Model string
	// Add other fields as needed
}

func fetchRecords(db *sql.DB, searchStr string, wg *sync.WaitGroup, resultsChan chan<- []Record) {
	defer wg.Done()

	query := "SELECT id, model FROM buybacks WHERE model LIKE ?"
	searchPattern := "%" + searchStr + "%"

	rows, err := db.Query(query, searchPattern)
	if err != nil {
		log.Println("Error querying database:", err)
		return
	}
	defer rows.Close()

	var records []Record

	for rows.Next() {
		var record Record
		if err := rows.Scan(&record.ID, &record.Model); err != nil {
			log.Println("Error scanning row:", err)
			return
		}
		records = append(records, record)
	}

	resultsChan <- records
}

func main() {
	dsn := "root:password@tcp(localhost:3306)/marketdata_test"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	searchStr := "Phone"

	const numWorkers = 4
	var wg sync.WaitGroup
	resultsChan := make(chan []Record, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go fetchRecords(db, searchStr, &wg, resultsChan)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var allRecords []Record
	for records := range resultsChan {
		allRecords = append(allRecords, records...)
	}

	for _, record := range allRecords {
		fmt.Printf("ID: %d, Model: %s\n", record.ID, record.Model)
	}
}
