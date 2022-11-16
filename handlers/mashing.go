package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/data"
)

type Mashing struct {
	l *log.Logger
}

func NewMashing(l *log.Logger) *Mashing {
	return &Mashing{l}
}

func (Mashing) GetProcedureToDo(rw http.ResponseWriter, r *http.Request) {
	mp := data.MashProcedure{
		RecipeId:       1,
		ProcedureCount: 2,
		MashProcedureList: []data.MashProcedureRecord{
			data.MashProcedureRecord{
				Temperature: 48,
				Holding:     16,
			},
			data.MashProcedureRecord{
				Temperature: 60,
				Holding:     12,
			},
		},
	}
	mashingBytes, err := json.MarshalIndent(mp, "", "\t")
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(mashingBytes)
}
