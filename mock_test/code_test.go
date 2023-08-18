package main

import (
	"sync"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSearchByTitle(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Set up the expected queries and responses for the mock
	mock.ExpectQuery("SELECT id, model FROM buybacks WHERE model LIKE ?").
		WithArgs("%Phone%").
		WillReturnRows(sqlmock.NewRows([]string{"id", "model"}).
			AddRow(1, "Phone Model 1").
			AddRow(2, "Phone Model 2"))

	// Create a WaitGroup
	var wg sync.WaitGroup

	// Set the number of goroutines to wait for
	wg.Add(1) // Assuming 1 asynchronous task

	// Call the function under test asynchronously
	go func() {
		defer wg.Done()
		searchStr := "Phone"
		_, err := searchByTitle(db, searchStr)
		assert.NoError(t, err, "searchByTitle returned an error")
	}()

	// Wait for the asynchronous task to finish
	wg.Wait()

	// Ensure all expected queries were executed
	assert.NoError(t, mock.ExpectationsWereMet(), "Unfulfilled expectations")
}
