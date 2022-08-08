package data

import (
	"encoding/json"
	"io"
	"time"
)

type Ingredient struct {
	Ingredient_Id   int                 `json:"ingredientId"`
	Ingredient_Name string              `json:"ingredientName" validate:"required"`
	Unit            string              `json:"ingredientStoreUnit" validate:"required,unit"`
	Quantity        float32             `json:"ingredientStoreQuantity" validate:"required"`
	CreatedAt       time.Time           `json:"ingrdientStoreCreatedAt"`
	Category        *IngredientCategory `json:"ingredientCategory"`
}

type IngredientVM struct {
	Ingredient_id       int     `json:"ingredientId"`
	Ingredient_name     string  `json:"ingredientName" validate:"required"`
	Ingredient_unit     string  `json:"ingredientStoreUnit" validate:"required,unit"`
	Ingredient_quantity float32 `json:"ingredientStoreQuantity" validate:"required"`
	Category_id         int     `json:"ingredientCategoryId"`
}

func (i *Ingredient) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

func (i *Ingredient) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i) // pass reference to myself, map json to struct
}
