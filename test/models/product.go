package models

import (
    "database/sql"
    "errors"
    "github.com/wpcodevo/golang-fiber-mysql/db"
    "fmt"
)

type Product struct {
    ID    int
    Name  string
    Price float64
}

func GetAllProducts() ([]Product, error) {
    fmt.Println("Starting Query...")
    query := "SELECT id, name, price FROM products"
    db := db.DB
    rows, err := db.Query(query)
    if err != nil {
        fmt.Println("Error executing query:", err)
        return nil, err
    }
    defer rows.Close()
    var products []Productt
    for rows.Next() {
        var product Product
        err := rows.Scan(&product.ID, &product.Name, &product.Price)
        if err != nil {
            fmt.Println("Error scanning row:", err)
            return nil, err
        }
        products = append(products, product)
    }
    fmt.Println("Query was successfully executed")
    return products, nil
}
func CreateProduct(product *Product) error {
    db := db.DB

    result, err := db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", product.Name, product.Price)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected != 1 {
        return errors.New("Failed to insert product")
    }

    return nil
}

func GetProductByID(id int) (*Product, error) {
    db := db.DB

    var product Product
    err := db.QueryRow("SELECT id, name, price FROM products WHERE id = ?", id).
        Scan(&product.ID, &product.Name, &product.Price)

    if err == sql.ErrNoRows {
        return nil, nil // No product found
    } else if err != nil {
        return nil, err
    }

    return &product, nil
}


func UpdateProduct(product *Product) error {
    db := db.DB

    result, err := db.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", product.Name, product.Price, product.ID)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected != 1 {
        return errors.New("Failed to update product")
    }

    return nil
}

func DeleteProduct(id int) error {
    db := db.DB

    result, err := db.Exec("DELETE FROM products WHERE id = ?", id)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected != 1 {
        return errors.New("Failed to delete product")
    }

    return nil
}
