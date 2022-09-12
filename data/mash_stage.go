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
}

type UpdateMashStageFormVM struct {
	Id          int     `json:"mashStageId"`
	StageTime   int64   `json:"stageTime"`
	Temperature float32 `json:"temperature"`
	PumpWork    bool    `json:"pumpWork"`
	RecipeId    int     `json:"recipeId"`
}

type MashStageFormVM struct {
	StageTime   int64   `json:"stageTime"`
	Temperature float32 `json:"temperature"`
	PumpWork    bool    `json:"pumpWork"`
	RecipeId    int     `json:"recipeId"`
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

func InsertMashStage(ms *MashStageFormVM) error {
	err, db := helpers.OpenConnection()
	if err != nil {
		db.Close()
		return err
	}
	defer db.Close()
	stm, err := db.Prepare("INSERT INTO mash_stage (temperature,stage_time,pump_work,recipe_id,created_at) VALUES($1,$2,$3,$4,NOW())")
	if err != nil {
		stm.Close()
		db.Close()
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(ms.Temperature, ms.StageTime, ms.PumpWork, ms.RecipeId)
	if err != nil {
		stm.Close()
		db.Close()
		return err
	}
	return nil
}

func DeleteMashStagesByRecipeId(recipeId int) error {
	err, db := helpers.OpenConnection()
	if err != nil {
		db.Close()
		return err
	}
	defer db.Close()
	smt, err := db.Prepare(`delete from mash_stage where recipe_id=$1;`)
	if err != nil {
		return err
	}
	defer smt.Close()
	if err != nil {
		return err
	}
	if _, err := smt.Exec(recipeId); err != nil {
		return err
	}
	return nil
}

func DeleteMashStage(id int) error {
	err, db := helpers.OpenConnection()
	if err != nil {
		db.Close()
		return err
	}
	defer db.Close()
	smt, err := db.Prepare(`delete from mash_stage where id=$1;`)
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

func UpdateMashStage(ms *UpdateMashStageFormVM) error {
	err, db := helpers.OpenConnection()
	if err != nil {
		return err
	}

	defer db.Close()
	smt, err := db.Prepare(`update mash_stage set temperature=$1, stage_time=$2, recipe_id=$3, pump_work=$4 where id=$5`)
	if err != nil {
		return err
	}
	defer smt.Close()

	if _, err := smt.Exec(ms.Temperature, ms.StageTime, ms.PumpWork, ms.RecipeId, ms.Id); err != nil {
		return err

	}
	return nil
}
