package handlers

import (
	"api-crud-categories/helpers"
	"api-crud-categories/models"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

var Categories = []models.Category{
	{ID: 1, Name: "Elektronik", Description: "Perangkat elektronik seperti HP, laptop, dan aksesoris"},
	{ID: 2, Name: "Fashion", Description: "Pakaian, sepatu, dan aksesoris fashion"},
	{ID: 3, Name: "Makanan", Description: "Produk makanan dan minuman"},
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		helpers.ResponseSuccess(w, Categories, http.StatusOK, "Categories retrieved successfully")
	} else {
		var newCategory models.Category
		err := json.NewDecoder(r.Body).Decode(&newCategory)
		if err != nil {
			helpers.ResponseError(w, http.StatusBadRequest, "Invalid request")
			return
		}

		newCategory.ID = len(Categories) + 1
		Categories = append(Categories, newCategory)

		helpers.ResponseSuccess(w, newCategory, http.StatusCreated, "Category created successfully")
	}
}

func CategoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getCategoryByID(w, r)
	} else if r.Method == "PUT" {
		updateCategory(w, r)
	} else if r.Method == "DELETE" {
		deleteCategory(w, r)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {

		helpers.ResponseError(w, http.StatusBadRequest, "Invalid Category ID")
		return
	}
	for _, c := range Categories {
		if c.ID == id {
			helpers.ResponseSuccess(w, c, http.StatusOK, "Category retrieved successfully")
			return
		}
	}
	helpers.ResponseError(w, http.StatusNotFound, "Category not found")
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {

		helpers.ResponseError(w, http.StatusBadRequest, "Invalid Category ID")
		return
	}
	for i, c := range Categories {
		if c.ID == id {
			Categories = append(Categories[:i], Categories[i+1:]...)

			helpers.ResponseSuccess(w, map[string]string{}, http.StatusOK, "Category deleted successfully")

			return
		}
	}
	helpers.ResponseError(w, http.StatusNotFound, "Category not found")
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {

		helpers.ResponseError(w, http.StatusBadRequest, "Invalid Category ID")
		return
	}
	var updateCategory models.Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {

		helpers.ResponseError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	for i := range Categories {
		if Categories[i].ID == id {
			updateCategory.ID = id
			Categories[i] = updateCategory
			helpers.ResponseSuccess(w, updateCategory, http.StatusOK, "Category updated successfully")
			return
		}
	}

}
