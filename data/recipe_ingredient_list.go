package data

import "time"

type RecipeIngredientList struct {
	Quantity   float32
	Unit       string
	CreatedAt  time.Time
	Recipe     *Recipe
	Ingredient *Ingredient
}
