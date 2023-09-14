package models

import (
    // "database/sql"
    // "errors"
    "github.com/wpcodevo/golang-fiber-mysql/db"
    "fmt"
)
type User struct {
    ID    int
    Name  string
    Age int
    Address string

}
func GetAllUsers() ([]User, error) {
    fmt.Println("Starting Query...")
    query := "SELECT id, name, age, address FROM users"
    db := db.DB
    rows, err := db.Query(query)
    if err != nil {
        fmt.Println("Error executing query:", err)
        return nil, err
    }
    defer rows.Close()
    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address)
        if err != nil {
            fmt.Println("Error scanning row:", err)
            return nil, err
        }
        users = append(users, user)
    }
    fmt.Println("Query was successfully executed")
    return users, nil
}
