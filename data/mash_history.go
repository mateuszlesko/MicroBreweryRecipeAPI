package data

import (
	"time"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/helpers"
)

type MashHistory struct {
	Id         int       `json:"mashHistoryId"`
	StartAt    time.Time `json:"startAt"`
	EndAt      time.Time `json:"endAt"`
	RecipeId   int       `json:"recipeId"`
	RecipeName string    `json:"recipeName"`
	MashTumId  int
}

type CurrentMashHistory struct {
	Id         int       `json:"mashHistoryId"`
	StartAt    time.Time `json:"startAt"`
	RecipeId   int       `json:"recipeId"`
	RecipeName string    `json:"recipeName"`
}

func CreateCurrentMash(mashId int, startAt time.Time, recipeId int, recipeName string) *CurrentMashHistory {
	currentMash := &CurrentMashHistory{}
	currentMash.Id = mashId
	currentMash.StartAt = startAt
	currentMash.RecipeId = recipeId
	currentMash.RecipeName = recipeName
	return currentMash
}

func InsertMashHistory(recipeId int) (int, error) {
	err, db := helpers.OpenConnection()
	if err != nil {
		return -1, err
	}
	var latestId int = 0
	err = db.QueryRow("Insert into mash_history(created_at,start_at,recipe_id,mash_tum_id) values(NOW(),NOW(),$1,1)  RETURNING id;").Scan(&latestId)
	if err != nil {
		return -1, err
	}

	return latestId, nil
}

func SelectMashHistory() ([]MashHistory, error) {
	err, db := helpers.OpenConnection()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("Select mash_history.id,mash_history.start_at,mash_history.end_at,mash_history.recipe_id, recipe.name from mash_history right join recipe on mash_history.recipe_id = recipe.id;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	mashHistory := []MashHistory{}

	var (
		mashId         int
		mashStartAt    time.Time
		mashEndAt      time.Time
		mashRecipeId   int
		mashRecipeName string
	)

	for rows.Next() {
		err = rows.Scan(&mashId, &mashStartAt, &mashEndAt, &mashRecipeId, &mashRecipeName)
		if err != nil {
			rows.Close()
			db.Close()
			return nil, err
		}
		mashHistory = append(mashHistory, MashHistory{mashId, mashStartAt, mashEndAt, mashRecipeId, mashRecipeName, 1})
	}
	return mashHistory, nil
}

func SelectCurrentMashings() ([]CurrentMashHistory, error) {
	err, db := helpers.OpenConnection()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("Select mash_history.id,mash_history.start_at,mash_history.recipe_id, recipe.name from mash_history right join recipe on mash_history.recipe_id = recipe.id where mash_history.end_at is null;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	mashHistory := []CurrentMashHistory{}

	var (
		mashId         int
		mashStartAt    time.Time
		mashRecipeId   int
		mashRecipeName string
	)

	for rows.Next() {
		err = rows.Scan(&mashId, &mashStartAt, &mashRecipeId, &mashRecipeName)
		if err != nil {
			rows.Close()
			db.Close()
			return nil, err
		}
		mashHistory = append(mashHistory, *CreateCurrentMash(mashId, mashStartAt, mashRecipeId, mashRecipeName))
	}
	return mashHistory, nil
}

func UpdateMashHistory(mashing_id int) error {
	err, db := helpers.OpenConnection()
	if err != nil {
		db.Close()
		return err
	}
	defer db.Close()
	smt, err := db.Prepare(`update mash_history set end_at=NOW() where id=$1`)
	if err != nil {
		smt.Close()
		db.Close()
		return err
	}
	if _, err := smt.Exec(mashing_id); err != nil {
		smt.Close()
		db.Close()
		return err
	}
	defer smt.Close()
	return nil
}
