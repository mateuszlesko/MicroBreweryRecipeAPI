package data

import (
	"encoding/json"
	"io"
	"time"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/helpers"
)

type Recipe struct {
	Id             int             `json:"recipeId"`
	RecipeName     string          `json:"recipeName"`
	CreatedAt      time.Time       `json:"recipeCreatedAt"`
	Denisty        float32         `json:"density"` //BLG
	IBU            float32         `json:"ibu"`
	RecipeCategory *RecipeCategory `json:"recipeCategory"`
}

type RecipeFormVM struct {
	Id               int     `json:"recipeId"`
	RecipeName       string  `json:"recipeName"`
	Denisty          float32 `json:"density"` //BLG
	IBU              float32 `json:"ibu"`
	RecipeCategoryId int     `json:"recipeCategoryId"`
}

func (c *Recipe) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func (c *Recipe) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(c) // pass reference to myself, map json to struct
}

func SelectRecipes() ([]Recipe, error) {
	err, db := helpers.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT recipe.id, recipe.name, recipe.density, recipe.ibu, recipe.created_at, recipe_category.id, recipe_category.name, recipe_category.created_at FROM recipe INNER JOIN recipe_category ON recipe.recipe_category_id = recipe_category.id;")
	if err != nil {
		db.Close()
		return nil, err
	}
	defer rows.Close()
	var recipes []Recipe
	var (
		recipeId                int
		recipeName              string
		recipeCreatedAt         time.Time
		recipeDensity           float32
		recipeIBU               float32
		recipeCategoryId        int
		recipeCategoryName      string
		recipeCategoryCreatedAt time.Time
	)
	for rows.Next() {
		if err := rows.Scan(&recipeId, &recipeName, &recipeDensity, &recipeIBU, &recipeCreatedAt, &recipeCategoryId, &recipeCategoryName, &recipeCategoryCreatedAt); err != nil {
			rows.Close()
			db.Close()
			return nil, err
		}
		recipes = append(recipes, Recipe{recipeId, recipeName, recipeCreatedAt, recipeDensity, recipeIBU, &RecipeCategory{recipeCategoryId, recipeCategoryName, recipeCategoryCreatedAt}})
	}
	return recipes, nil
}

func InsertRecipe(rfvm *RecipeFormVM) error {
	err, db := helpers.OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO RECIPE (name,density,ibu,created_at,recipe_category_id) VALUES($1,$2,$3,NOW(),$4);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(rfvm.RecipeName, rfvm.Denisty, rfvm.IBU, rfvm.RecipeCategoryId)
	if err != nil {
		return err
	}
	return nil
}

func UpdateRecipe(rc RecipeFormVM) error {
	err, db := helpers.OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	smt, err := db.Prepare(`update recipe set name=$1 recipe_category_id=$2 density=$3 ibu=$4 where id=$5`)
	if err != nil {
		return err
	}
	defer smt.Close()

	if _, err := smt.Exec(rc.RecipeName, rc.RecipeCategoryId, rc.Denisty, rc.IBU, rc.Id); err != nil {
		return err
	}
	return nil
}

func DeleteRecipe(recipeId int) error {
	err, db := helpers.OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	smt, err := db.Prepare(`delete from recipe_category where id=$1;`)
	if err != nil {
		return err
	}
	defer smt.Close()

	if _, err := smt.Exec(recipeId); err != nil {
		return err
	}
	return nil
}
