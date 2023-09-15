package controllers

import (
    "html/template"
    "net/http"
    "strconv"
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
func UserCreate(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        name := r.FormValue("name")
        ageStr := r.FormValue("age")
        address := r.FormValue("address")

        age, err := strconv.Atoi(ageStr)
        if err != nil {
            http.Error(w, "Invalid age", http.StatusBadRequest)
            return
        }
        user := models.User{
            Name:    name,
            Age:     age,
            Address: address,
        }
        err = models.CreateUser(&user)
        if err != nil {
            http.Error(w, "Error creating user", http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, "/users", http.StatusSeeOther)
        return
    }
    tpl := template.Must(template.ParseFiles("templates/user/create.html"))
    tpl.Execute(w, nil)
}
func UserEdit(w http.ResponseWriter, r *http.Request) {
    userIDStr := r.URL.Query().Get("id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    user, err := models.GetUserByID(userID)
    if err != nil {
        http.Error(w, "Error fetching user", http.StatusInternalServerError)
        return
    }
    if r.Method == http.MethodPost {
        name := r.FormValue("name")
        ageStr := r.FormValue("age")
        address := r.FormValue("address")

        age, err := strconv.Atoi(ageStr)
        if err != nil {
            http.Error(w, "Invalid age", http.StatusBadRequest)
            return
        }
        user.Name = name
        user.Age = age
        user.Address = address
        err = models.UpdateUser(user)
        if err != nil {
            http.Error(w, "Error updating user", http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, "/users", http.StatusSeeOther)
        return
    }
    tpl := template.Must(template.ParseFiles("templates/user/edit.html"))
    tpl.Execute(w, user)
}
func UserDelete(w http.ResponseWriter, r *http.Request) {
    userIDStr := r.URL.Query().Get("id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    err = models.DeleteUser(userID)
    if err != nil {
        http.Error(w, "Error deleting user", http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/users", http.StatusSeeOther)
}
func UserView(w http.ResponseWriter, r *http.Request) {
    userIDStr := r.URL.Query().Get("id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    user, err := models.GetUserByID(userID)
    if err != nil {
        http.Error(w, "Error fetching user", http.StatusInternalServerError)
        return
    }
    tpl := template.Must(template.ParseFiles("templates/user/view.html"))
    tpl.Execute(w, user)
}
