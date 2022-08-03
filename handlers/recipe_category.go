package handlers

import (
	"log"
	"net/http"
)

type RecipeCategory struct {
	l *log.Logger
}

func NewRecipeCategory(logger *log.Logger) *RecipeCategory {
	return &RecipeCategory{logger}
}

func (rch *RecipeCategory) GetRecipeCategories(rw http.ResponseWriter, r *http.Request) {

}

func (rch *RecipeCategory) PostRecipeCategory(rw http.ResponseWriter, r *http.Request) {

}

func (rch *RecipeCategory) PutRecipeCategory(rw http.ResponseWriter, r *http.Request) {

}

func (rch *RecipeCategory) DeleteRecipeCategory(rw http.ResponseWriter, r *http.Request) {

}
