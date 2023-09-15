package main

import (
    "net/http"
    "github.com/wpcodevo/golang-fiber-mysql/controllers"
    "github.com/wpcodevo/golang-fiber-mysql/db"
)

func main() {
    db.InitDB()

    http.HandleFunc("/products", controllers.ProductIndex)
    http.HandleFunc("/products/create", controllers.ProductCreate)
    http.HandleFunc("/products/edit", controllers.ProductEdit)
    http.HandleFunc("/products/view", controllers.ProductView)
    http.HandleFunc("/products/delete", controllers.ProductDelete)

    http.HandleFunc("/users", controllers.UserIndex)
    http.HandleFunc("/users/create", controllers.UserCreate)
    http.HandleFunc("/users/edit", controllers.UserEdit)
    http.HandleFunc("/users/view", controllers.UserView)
    http.HandleFunc("/users/delete", controllers.UserDelete)

    http.HandleFunc("/orders", controllers.OrderIndex)
    http.HandleFunc("/orders/create", controllers.OrderCreate)
    http.HandleFunc("/orders/edit", controllers.OrderEdit)
    http.HandleFunc("/orders/delete", controllers.OrderDelete)

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    http.ListenAndServe(":8081", nil)
}
