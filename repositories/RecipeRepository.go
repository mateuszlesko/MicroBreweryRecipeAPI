package repositories

import (
<<<<<<< HEAD
	"log"

=======
>>>>>>> 9880dbd9b2b07cf6782982fc2e407c749bf51fde
	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/data"
)

type RecipeRepository struct {
}

func CreateRecipeRepository() *RecipeRepository {
	return &RecipeRepository{}
}

func (rr *RecipeRepository) GetFullRecipeData(recipeId int) (data.RecipeFullData, error) {
<<<<<<< HEAD
	recipe, err := data.SelectRecipeById(recipeId)
	if err != nil {
		log.Panic()
		return data.RecipeFullData{}, err
	}
	mashstages, err := data.SelectMashStagesByRecipeId(recipeId)
	if err != nil {
		return data.RecipeFullData{}, err
	}
	var recipeFullData data.RecipeFullData = data.RecipeFullData{}
	recipeFullData.Recipe = recipe
	recipeFullData.MashStages = mashstages
=======
	recipe, _ := data.SelectRecipeById(recipeId)
	mashstages, _ := data.SelectMashStagesByRecipeId(recipeId)
	ingredientList, _ := data.SelectRecipeIngredientList(recipeId)
	var recipeFullData data.RecipeFullData = data.RecipeFullData{}
	recipeFullData.Recipe = recipe
	recipeFullData.MashStages = mashstages
	recipeFullData.RecipeIngredientList = ingredientList
>>>>>>> 9880dbd9b2b07cf6782982fc2e407c749bf51fde

	return recipeFullData, nil
}
