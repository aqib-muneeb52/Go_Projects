package main

import (
    "net/http"
    "github.com/wpcodevo/golang-fiber-mysql/controllers"
    "github.com/wpcodevo/golang-fiber-mysql/db" // Import your database package
)

func main() {
    db.InitDB()
    http.HandleFunc("/products", controllers.ProductIndex)
    http.HandleFunc("/products/create", controllers.ProductCreate)
    http.HandleFunc("/products/edit", controllers.ProductEdit)
    http.HandleFunc("/products/view", controllers.ProductView)
    http.HandleFunc("/products/delete", controllers.ProductDelete)
    http.HandleFunc("/users", controllers.UserIndex)

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.ListenAndServe(":8081", nil)
}
