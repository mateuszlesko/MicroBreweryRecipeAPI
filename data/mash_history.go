package data

import (
	"time"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/helpers"
)

type MashHistory struct {
	Id        int       `json:"mashHistoryId"`
	StartAt   time.Time `json:"startAt"`
	EndAt     time.Time `json:"endAt"`
	Recipe    Recipe    `json:"recipe"`
	MashTumId int
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

// func SelectMashHistory() ([]MashHistory, error) {
// 	err, db := helpers.OpenConnection()
// 	if err != nil {
// 		return nil, err
// 	}
// 	smt, err :=
// 	return []MashHistory{}, nil
// }
