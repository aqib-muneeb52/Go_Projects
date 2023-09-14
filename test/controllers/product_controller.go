package controllers

import (
    "html/template"
    "net/http"
    "strconv"
    "github.com/wpcodevo/golang-fiber-mysql/models"
)

func ProductIndex(w http.ResponseWriter, r *http.Request) {
    products, err := models.GetAllProducts()
    if err != nil {
        http.Error(w, "Error fetching products", http.StatusInternalServerError)
        return
    }
    tpl := template.Must(template.ParseFiles("templates/product/index.html"))
    tpl.Execute(w, products)
}

func ProductCreate(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        name := r.FormValue("name")
        priceStr := r.FormValue("price")
        price, err := strconv.ParseFloat(priceStr, 64)
        if err != nil {
            http.Error(w, "Invalid price", http.StatusBadRequest)
            return
        }
        product := &models.Product{
            Name:  name,
            Price: price,
        }
        err = models.CreateProduct(product)
        if err != nil {
            http.Error(w, "Error creating product", http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, "/products", http.StatusSeeOther)
        return
    }
    tpl := template.Must(template.ParseFiles("templates/product/create.html"))
    tpl.Execute(w, nil)
}

func ProductEdit(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        idStr := r.FormValue("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.Error(w, "Invalid product ID", http.StatusBadRequest)
            return
        }
        name := r.FormValue("name")
        priceStr := r.FormValue("price")
        price, err := strconv.ParseFloat(priceStr, 64)
        if err != nil {
            http.Error(w, "Invalid price", http.StatusBadRequest)
            return
        }
        product := &models.Product{
            ID:    id,
            Name:  name,
            Price: price,
        }
        err = models.UpdateProduct(product)
        if err != nil {
            http.Error(w, "Error updating product", http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, "/products", http.StatusSeeOther)
        return
    }
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }
    product, err := models.GetProductByID(id)
    if err != nil {
        http.Error(w, "Error fetching product", http.StatusInternalServerError)
        return
    }
    tpl := template.Must(template.ParseFiles("templates/product/edit.html"))
    tpl.Execute(w, product)
}

func ProductView(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }
    product, err := models.GetProductByID(id)
    if err != nil {
        http.Error(w, "Error fetching product", http.StatusInternalServerError)
        return
    }
    tpl := template.Must(template.ParseFiles("templates/product/view.html"))
    tpl.Execute(w, product)
}

func ProductDelete(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }
    err = models.DeleteProduct(id)
    if err != nil {
        http.Error(w, "Error deleting product", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/products", http.StatusSeeOther)
}
