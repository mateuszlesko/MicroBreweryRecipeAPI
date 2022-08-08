package data

import (
	"encoding/json"
	"io"
	"time"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/helpers"
)

type RecipeCategory struct {
	Id         int       `json:"recipeCategoryId"`
	Name       string    `json:"recipeCategoryName"`
	Created_at time.Time `json:"recipeCategoryCreatedAt"`
}

type RecipeCategoryFormVM struct {
	Id   int    `json:"recipeCategoryId"`
	Name string `json:"recipeCategoryName"`
}

func (rc *RecipeCategory) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(rc)
}

func (rc *RecipeCategory) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(rc)
}

func ToJSON(rcs []RecipeCategory, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(rcs)
}

func SelectRecipeCategories() ([]RecipeCategory, error) {
	err, db := helpers.OpenConnection()
	if err != nil {
		db.Close()
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("Select id,name,created_at from recipe_category")
	if err != nil {
		rows.Close()
		db.Close()
		return nil, err
	}
	defer rows.Close()
	var categories []RecipeCategory
	for rows.Next() {
		var category RecipeCategory
		if err := rows.Scan(&category.Id, &category.Name, &category.Created_at); err != nil {
			rows.Close()
			db.Close()
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func InsertRecipeCategory(name string) error {
	err, db := helpers.OpenConnection()
	if err != nil {
		db.Close()
		return err
	}
	defer db.Close()
	stm, err := db.Prepare("INSERT INTO recipe_category (name,created_at) VALUES($1,NOW())")
	if err != nil {
		stm.Close()
		db.Close()
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(name)
	if err != nil {
		stm.Close()
		db.Close()
		return err
	}
	return nil
}

func UpdateRecipeCategory(rc RecipeCategory) (*RecipeCategory, error) {
	err, db := helpers.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	smt, err := db.Prepare(`update recipe_category set name=$1 where id=$2`)
	if err != nil {
		return nil, err
	}
	defer smt.Close()

	if _, err := smt.Exec(rc.Name, rc.Id); err != nil {
		return nil, err
	}

	return &rc, nil
}

func DeleteRecipeCategory(id int) error {
	err, db := helpers.OpenConnection()
	if err != nil {
		db.Close()
		return err
	}
	defer db.Close()
	smt, err := db.Prepare("update recipe set recipe_category_id=NULL where recipe_category_id = $1")
	if err != nil {
		smt.Close()
		db.Close()
		return err
	}
	defer smt.Close()
	if _, err = smt.Exec(id); err != nil {
		return err
	}
	smt, err = db.Prepare(`delete from recipe_category where id=$1;`)
	if err != nil {
		return err
	}
	defer smt.Close()
	if err != nil {
		return err
	}
	if _, err := smt.Exec(id); err != nil {
		return err
	}

	return nil
}
