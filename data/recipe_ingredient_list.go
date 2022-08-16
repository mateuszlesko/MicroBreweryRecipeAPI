package data

import (
	"encoding/json"
	"io"
	"time"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/helpers"
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

func SelectRecipeIngredientList(recipeId int) ([]RecipeIngredientList, error) {
	//;
	err, db := helpers.OpenConnection()
	if err != nil {
		db.Close()
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("select recipe_ingredient_list.quantity,recipe_ingredient_list.unit,ingredient.id, ingredient.ingredient_name from recipe_ingredient_list inner join ingredient on recipe_ingredient_list.ingredient_id = ingredient.id where recipe_ingredient_list.recipe_id = $1;", recipeId)
	if err != nil {
		rows.Close()
		db.Close()
		return nil, err
	}
	defer rows.Close()
	var (
		ingredientList  []RecipeIngredientList
		ingredient_id   int
		ingredient_name string
	)
	for rows.Next() {
		var il RecipeIngredientList
		if err := rows.Scan(&il.Quantity, &il.Unit, &ingredient_id, &ingredient_name); err != nil {
			rows.Close()
			db.Close()
			return nil, err
		}
		il.Ingredient = &Ingredient{Ingredient_Id: ingredient_id, Ingredient_Name: ingredient_name}
		ingredientList = append(ingredientList, il)
	}
	return ingredientList, nil
}
