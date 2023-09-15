package models

import (
    // "database/sql"
    "github.com/wpcodevo/golang-fiber-mysql/db"
)

type Order struct {
    ID          int
    Total       float64
    ProductID   int
    UserID      int
    ProductName string
    UserName    string
}

func CreateOrder(order *Order) error {
    db := db.DB

    _, err := db.Exec("INSERT INTO orders (total, product_id, user_id) VALUES (?, ?, ?)", order.Total, order.ProductID, order.UserID)
    if err != nil {
        return err
    }

    return nil
}

func GetOrderByID(id int) (*Order, error) {
    db := db.DB

    row := db.QueryRow("SELECT o.id, o.total, o.product_id, o.user_id, p.name AS product_name, u.name AS user_name FROM orders o INNER JOIN products p ON o.product_id = p.id INNER JOIN users u ON o.user_id = u.id WHERE o.id = ?", id)

    var order Order
    err := row.Scan(&order.ID, &order.Total, &order.ProductID, &order.UserID, &order.ProductName, &order.UserName)
    if err != nil {
        return nil, err
    }

    return &order, nil
}

func GetAllOrders() ([]Order, error) {
    db := db.DB

    rows, err := db.Query("SELECT o.id, o.total, o.product_id, o.user_id, p.name AS product_name, u.name AS user_name FROM orders o INNER JOIN products p ON o.product_id = p.id INNER JOIN users u ON o.user_id = u.id")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var orders []Order

    for rows.Next() {
        var order Order
        err := rows.Scan(&order.ID, &order.Total, &order.ProductID, &order.UserID, &order.ProductName, &order.UserName)
        if err != nil {
            return nil, err
        }
        orders = append(orders, order)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return orders, nil
}

func UpdateOrder(order *Order) error {
    db := db.DB

    _, err := db.Exec("UPDATE orders SET total = ?, product_id = ? WHERE id = ?", order.Total, order.ProductID, order.ID)
    if err != nil {
        return err
    }

    return nil
}

func DeleteOrder(id int) error {
    db := db.DB

    _, err := db.Exec("DELETE FROM orders WHERE id = ?", id)
    if err != nil {
        return err
    }

    return nil
}
