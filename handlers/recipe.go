package handlers

import (
	"log"
	"net/http"
)

type Recipe struct {
	l *log.Logger
}

func NewRecipe(l *log.Logger) *Recipe {
	return &Recipe{l}
}

func (rh *Recipe) GetRecipes(rw http.ResponseWriter, r *http.Request) {

}

func (rh *Recipe) PostRecipe(rw http.ResponseWriter, r *http.Request) {

}

func (rh *Recipe) PutRecipe(rw http.ResponseWriter, r *http.Request) {

}

func (rh *Recipe) DeleteRecipe(rw http.ResponseWriter, r *http.Request) {

}
