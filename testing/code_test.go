package main

import (
    "database/sql"
    "testing"

    "github.com/DATA-DOG/go-sqlmock"
)

func TestSearchByTitle(t *testing.T) {
	// Existing test case (unchanged)

	t.Run("No results found", func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
					t.Fatalf("Failed to create mock: %v", err)
			}
			defer db.Close()

			mock.ExpectQuery("SELECT id, model").
					WithArgs("%NonExistent%").
					WillReturnRows(sqlmock.NewRows([]string{"id", "model"}))

			records, err := searchByTitle(db, "NonExistent")
			if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
			}

			if len(records) != 0 {
					t.Errorf("Expected 0 records, but got %d records", len(records))
			}
	})

	t.Run("Database query error", func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
					t.Fatalf("Failed to create mock: %v", err)
			}
			defer db.Close()

			mock.ExpectQuery("SELECT id, model").
					WithArgs("%Error%").
					WillReturnError(sql.ErrConnDone)

			_, err = searchByTitle(db, "Error")
			if err == nil {
					t.Error("Expected an error, but got no error")
			}
	})

	// Add more test cases as needed
}

