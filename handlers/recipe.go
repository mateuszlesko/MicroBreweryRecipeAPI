package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/data"
	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/repositories"
)

type Recipe struct {
	l *log.Logger
}

func NewRecipe(l *log.Logger) *Recipe {
	return &Recipe{l}
}

func (rh *Recipe) GetRecipes(rw http.ResponseWriter, r *http.Request) {
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

func (rh *Recipe) GetRecipeById(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
<<<<<<< HEAD
	fmt.Println(id)
	if err != nil {
		http.Error(rw, "unable to get argument", http.StatusBadRequest)
	}
	//recipe, err := data.SelectRecipeById(id)
=======
	if err != nil {
		http.Error(rw, "unable to get argument", http.StatusBadRequest)
	}
>>>>>>> 9880dbd9b2b07cf6782982fc2e407c749bf51fde
	recipeRepo := repositories.CreateRecipeRepository()
	recipe, err := recipeRepo.GetFullRecipeData(id)
	if err != nil {
		http.Error(rw, "unable to get data", http.StatusBadRequest)
	}
	recipeBytes, err := json.MarshalIndent(recipe, "", "\t")
	if err != nil {
		http.Error(rw, "unable to parse data", http.StatusBadRequest)
	}
	rw.Write(recipeBytes)
}

func (rh *Recipe) PostRecipe(rw http.ResponseWriter, r *http.Request) {

}

func (rh *Recipe) PutRecipe(rw http.ResponseWriter, r *http.Request) {

}

func (rh *Recipe) DeleteRecipe(rw http.ResponseWriter, r *http.Request) {

}
