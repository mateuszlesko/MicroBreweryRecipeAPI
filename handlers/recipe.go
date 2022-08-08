package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/data"
)

type Recipe struct {
	l *log.Logger
}

func NewRecipe(l *log.Logger) *Recipe {
	return &Recipe{l}
}

func (rh *Recipe) GetRecipes(rw http.ResponseWriter, r *http.Request) {
	fmt.Printf("XD")
	rl, err := data.SelectRecipes()
	if err != nil {
		http.Error(rw, "Unable to get data", http.StatusBadRequest)
	}
	recipesBytes, err := json.MarshalIndent(rl, "", "\t")
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(recipesBytes)
}

func (rh *Recipe) PostRecipe(rw http.ResponseWriter, r *http.Request) {

}

func (rh *Recipe) PutRecipe(rw http.ResponseWriter, r *http.Request) {

}

func (rh *Recipe) DeleteRecipe(rw http.ResponseWriter, r *http.Request) {

}
