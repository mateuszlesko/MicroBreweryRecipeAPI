package data

import "time"

type Ingredient struct {
	Ingredient_Id   int                 `json:"id"`
	Ingredient_Name string              `json:"name" validate:"required"`
	Unit            string              `json:"unit" validate:"required,unit"`
	Quantity        float32             `json:"quantity" validate:"required"`
	CreatedAt       time.Time           `json:"created_at"`
	Category        *IngredientCategory `json:"category"`
}

type IngredientVM struct {
	Ingredient_id          int     `json:"id"`
	Ingredient_name        string  `json:"name" validate:"required"`
	Ingredient_unit        string  `json:"unit" validate:"required,unit"`
	Ingredient_quantity    float32 `json:"quantity" validate:"required"`
	Ingredient_description string  `json:"desc"`
	Category_id            int     `json:"category"`
}
