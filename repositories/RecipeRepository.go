package repositories

import (
	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/data"
)

type RecipeRepository struct {
}

func CreateRecipeRepository() *RecipeRepository {
	return &RecipeRepository{}
}

func (rr *RecipeRepository) GetFullRecipeData(recipeId int) (data.RecipeFullData, error) {
	recipe, _ := data.SelectRecipeById(recipeId)
	mashstages, _ := data.SelectMashStagesByRecipeId(recipeId)
	ingredientList, _ := data.SelectRecipeIngredientList(recipeId)
	var recipeFullData data.RecipeFullData = data.RecipeFullData{}
	recipeFullData.Recipe = recipe
	recipeFullData.MashStages = mashstages
	recipeFullData.RecipeIngredientList = ingredientList

	return recipeFullData, nil
}
