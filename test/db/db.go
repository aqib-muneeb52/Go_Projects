package db

import (
    "database/sql"
    "log"
		"fmt"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	  fmt.Println("Initializing database...")
    var err error
    DB, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/go")
    if err != nil {
        log.Fatal(err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal(err)
    }
}
