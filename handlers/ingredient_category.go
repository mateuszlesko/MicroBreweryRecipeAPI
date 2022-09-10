package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/data"
)

type Category struct {
	l *log.Logger
}

type KeyCategory struct{}

func NewIngredientCategory(l *log.Logger) *Category {
	return &Category{l}
}

//get
func (c *Category) GetIngredientCategories(rw http.ResponseWriter, r *http.Request) {
	cl, err := data.SelectIngredientCategories()
	if err != nil {
		http.Error(rw, "unable to query", http.StatusUnprocessableEntity)
		return
	}
	categoriesBytes, err := json.MarshalIndent(cl, "", "\t")
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(categoriesBytes)
}

//get
func (c *Category) GetIngredientCategory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil && id != 0 {
		http.Error(rw, "unable to decode value", http.StatusBadRequest)
		return
	}
	var category *data.IngredientCategory
	category, err = data.SelectIngredientCategoryWhereID(id)
	if err != nil {
		c.l.Panic(err)
	}
	categoryBytes, err := json.Marshal(category)
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(categoryBytes)
}

//delete
func (c *Category) DeleteIngredientCategory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to decode json", http.StatusBadRequest)
		return
	}
	_, err = data.SelectIngredientCategoryWhereID(id)
	if err != nil {
		http.Error(rw, "no object corresponds in db", http.StatusBadRequest)
		return
	}
	err = data.DeleteCategory(id)
	if err != nil {
		http.Error(rw, "delete was not executed", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte("1"))
}

func (c *Category) UpdateIngredientCategory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to decode json", http.StatusBadRequest)
		return
	}
	_, err = data.SelectIngredientCategoryWhereID(id)
	if err != nil {
		http.Error(rw, "no object corresponds in db", http.StatusBadGateway)
		return
	}
	category := data.IngredientCategory{}
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(rw, "not enable to decode json", http.StatusBadGateway)
	}
	_, err = data.UpdateIngredientCategory(category)
	if err != nil {
		http.Error(rw, "unable to update", http.StatusUnprocessableEntity)
		return
	}
	categoryBytes, err := json.Marshal(category)
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(categoryBytes)
}

func (c Category) PostIngredientCategory(rw http.ResponseWriter, r *http.Request) {
	var category data.IngredientCategory
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(rw, "can not decode value from body", http.StatusBadRequest)
		return
	}
	err = data.InsertIngredientCategory(category.Category_name)
	if err != nil {
		http.Error(rw, "can not add row", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write([]byte(fmt.Sprintf("%d", 1)))
}
