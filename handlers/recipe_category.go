package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/data"
)

type RecipeCategory struct {
	l *log.Logger
}

func NewRecipeCategory(logger *log.Logger) *RecipeCategory {
	return &RecipeCategory{logger}
}

func (rch *RecipeCategory) GetRecipeCategories(rw http.ResponseWriter, r *http.Request) {
	rcl, err := data.SelectRecipeCategories()
	if err != nil {
		http.Error(rw, "Unable to get data", http.StatusBadRequest)
	}
	categoriesBytes, err := json.MarshalIndent(rcl, "", "\t")
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(categoriesBytes)
}

func (rch *RecipeCategory) PostRecipeCategory(rw http.ResponseWriter, r *http.Request) {
	var category data.RecipeCategory
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(rw, "can not decode value", http.StatusBadRequest)
	}
	err = data.InsertRecipeCategory(category.Name)
	if err != nil {
		http.Error(rw, "can not create value", http.StatusBadRequest)
	}

}

func (rch *RecipeCategory) PutRecipeCategory(rw http.ResponseWriter, r *http.Request) {
	var category data.RecipeCategory
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(rw, "can not decode value", http.StatusBadRequest)
	}
	_, err = data.UpdateRecipeCategory(category)
	if err != nil {
		http.Error(rw, "can not update value", http.StatusBadRequest)
	}
}

func (rch *RecipeCategory) DeleteRecipeCategory(rw http.ResponseWriter, r *http.Request) {

}
