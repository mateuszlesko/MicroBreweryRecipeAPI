package repositories

import (
	"log"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/data"
)

type RecipeRepository struct {
}

func CreateRecipeRepository() *RecipeRepository {
	return &RecipeRepository{}
}

func (rr *RecipeRepository) GetFullRecipeData(recipeId int) (data.RecipeFullData, error) {
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

	return recipeFullData, nil
}
