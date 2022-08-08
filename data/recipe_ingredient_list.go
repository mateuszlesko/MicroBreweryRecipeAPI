package data

import (
	"encoding/json"
	"io"
	"time"
)

type RecipeIngredientList struct {
	Quantity   float32     `json:"recipeIngredientQuantity"`
	Unit       string      `json:"recipeIngredientUnit"`
	CreatedAt  time.Time   `json:"recipeIngredientId"`
	Recipe     *Recipe     `json:"recipe"`
	Ingredient *Ingredient `json:"ingredient"`
}

type RecipeIngredientListFormVM struct {
	Quantity   float32 `json:"recipeIngredientQuantity"`
	Unit       string  `json:"recipeIngredientUnit"`
	Recipe     int     `json:"recipeId"`
	Ingredient int     `json:"ingredientId"`
}

func (c *RecipeIngredientList) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func (c *RecipeIngredientList) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(c) // pass reference to myself, map json to struct
}
