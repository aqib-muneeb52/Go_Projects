package controllers

import (
    "html/template"
    "net/http"
    "github.com/wpcodevo/golang-fiber-mysql/models"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
    products, err := models.GetAllUsers()
    if err != nil {
        http.Error(w, "Error fetching users", http.StatusInternalServerError)
        return
    }
    tpl := template.Must(template.ParseFiles("templates/user/index.html"))
    tpl.Execute(w, products)
}
