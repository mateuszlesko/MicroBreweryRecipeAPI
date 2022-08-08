package data

import (
	"encoding/json"
	"io"
	"time"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/helpers"
)

type MashStage struct {
	Id          int       `json:"meshStageId"`
	StageTime   int64     `json:"stageTime"` //ms
	Temperature float32   `json:"temperature"`
	PumpWork    bool      `json:"pumpWork"`
	CreatedAt   time.Time `json:"createdAt"`
	Recipe      *Recipe   `json:"recipe"`
}
type MashStageFormVM struct {
	StageTime   int64   `json:"stageTime"`
	Temperature float32 `json:"temperature"`
	PumpWork    bool    `json:"pumpWork"`
	Recipe_id   int     `json:"recipeId"`
}

func (c *MashStage) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func (c *MashStage) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(c) // pass reference to myself, map json to struct
}

func SelectMashStagesByRecipeId(recipeId int) ([]MashStage, error) {
	err, db := helpers.OpenConnection()
	if err != nil {
		db.Close()
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT mash_stage.temperature,mash_stage.stage_time,mash_stage.pump_work FROM mash_stage WHERE mash_stage.recipe_id = $1;", recipeId)
	if err != nil {
		rows.Close()
		db.Close()
		return nil, err
	}
	defer rows.Close()
	var mashStages []MashStage
	for rows.Next() {
		var ms MashStage
		if err := rows.Scan(&ms.Temperature, &ms.StageTime, &ms.PumpWork); err != nil {
			rows.Close()
			db.Close()
			return nil, err
		}
		mashStages = append(mashStages, ms)
	}
	return mashStages, nil
}
