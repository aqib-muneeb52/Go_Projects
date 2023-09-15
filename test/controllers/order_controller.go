package controllers

import (
    "html/template"
    "net/http"
    "strconv"
    "github.com/wpcodevo/golang-fiber-mysql/models"
)

func OrderIndex(w http.ResponseWriter, r *http.Request) {
	orders, err := models.GetAllOrders()
	if err != nil {
			http.Error(w, "Error fetching orders", http.StatusInternalServerError)
			return
	}
	tpl := template.Must(template.ParseFiles("templates/order/index.html"))
	tpl.Execute(w, orders)
}

func OrderCreate(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        totalStr := r.FormValue("total")
        total, err := strconv.ParseFloat(totalStr, 64)
        if err != nil {
            http.Error(w, "Invalid total", http.StatusBadRequest)
            return
        }
        productIDStr := r.FormValue("productID")
        productID, err := strconv.Atoi(productIDStr)
        if err != nil {
            http.Error(w, "Invalid product ID", http.StatusBadRequest)
            return
        }
        userIDStr := r.FormValue("userID")
        userID, err := strconv.Atoi(userIDStr)
        if err != nil {
            http.Error(w, "Invalid user ID", http.StatusBadRequest)
            return
        }
        order := &models.Order{
            Total:     total,
            ProductID: productID,
            UserID:    userID,
        }
        err = models.CreateOrder(order)
        if err != nil {
            http.Error(w, "Error creating order", http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, "/orders", http.StatusSeeOther)
        return
    }
    products, err := models.GetAllProducts()
    if err != nil {
        http.Error(w, "Error fetching products", http.StatusInternalServerError)
        return
    }
    users, err := models.GetAllUsers()
    if err != nil {
        http.Error(w, "Error fetching users", http.StatusInternalServerError)
        return
    }
    tpl := template.Must(template.ParseFiles("templates/order/create.html"))
    tpl.Execute(w, struct {
        Products []models.Product
        Users    []models.User
    }{
        Products: products,
        Users:    users,
    })
}

func OrderEdit(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid order ID", http.StatusBadRequest)
        return
    }
    order, err := models.GetOrderByID(id)
    if err != nil {
        http.Error(w, "Error fetching order", http.StatusInternalServerError)
        return
    }
    if r.Method == http.MethodPost {
        totalStr := r.FormValue("total")
        total, err := strconv.ParseFloat(totalStr, 64)
        if err != nil {
            http.Error(w, "Invalid total", http.StatusBadRequest)
            return
        }
        productIDStr := r.FormValue("productID")
        productID, err := strconv.Atoi(productIDStr)
        if err != nil {
            http.Error(w, "Invalid product ID", http.StatusBadRequest)
            return
        }
        order.Total = total
        order.ProductID = productID
        err = models.UpdateOrder(order)
        if err != nil {
            http.Error(w, "Error updating order", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/orders", http.StatusSeeOther)
        return
    }
    products, err := models.GetAllProducts()
    if err != nil {
        http.Error(w, "Error fetching products", http.StatusInternalServerError)
        return
    }
    tpl := template.Must(template.ParseFiles("templates/order/edit.html"))
    tpl.Execute(w, struct {
        Order    *models.Order
        Products []models.Product
    }{
        Order:    order,
        Products: products,
    })
}

func OrderDelete(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid order ID", http.StatusBadRequest)
        return
    }
    err = models.DeleteOrder(id)
    if err != nil {
        http.Error(w, "Error deleting order", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/orders", http.StatusSeeOther)
}
