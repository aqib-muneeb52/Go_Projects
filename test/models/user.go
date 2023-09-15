package models

import (
    "github.com/wpcodevo/golang-fiber-mysql/db"
    "fmt"
)

// ...
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

func CreateUser(user *User) error {
    db := db.DB

    _, err := db.Exec("INSERT INTO users (name, age, address) VALUES (?, ?, ?)", user.Name, user.Age, user.Address)
    if err != nil {
        return err
    }

    return nil
}

func UpdateUser(user *User) error {
    db := db.DB

    _, err := db.Exec("UPDATE users SET name=?, age=?, address=? WHERE id=?", user.Name, user.Age, user.Address, user.ID)
    if err != nil {
        return err
    }

    return nil
}

func DeleteUser(userID int) error {
    db := db.DB

    _, err := db.Exec("DELETE FROM users WHERE id=?", userID)
    if err != nil {
        return err
    }

    return nil
}

func GetUserByID(userID int) (*User, error) {
    db := db.DB

    var user User
    err := db.QueryRow("SELECT id, name, age, address FROM users WHERE id=?", userID).Scan(&user.ID, &user.Name, &user.Age, &user.Address)
    if err != nil {
        return nil, err
    }

    return &user, nil
}
