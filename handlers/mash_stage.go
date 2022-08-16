package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/data"
)

type MashStage struct {
	l *log.Logger
}

func NewMashStage(l *log.Logger) *MashStage {
	return &MashStage{l}
}

func (mh *MashStage) GetMashStageByRecipeId(rw http.ResponseWriter, r *http.Request) {

	recipeId, err := strconv.Atoi(r.URL.Query().Get("recipeId"))
	if err != nil || recipeId == 0 {
		log.Panic()
		http.Error(rw, "unable to parse value", http.StatusBadRequest)
	}
	msl, err := data.SelectMashStagesByRecipeId(recipeId)
	if err != nil {
		http.Error(rw, "unable to get data", http.StatusUnprocessableEntity)
		return
	}
	mashBytes, err := json.Marshal(msl)
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(mashBytes)
}

func (mh *MashStage) PostMashStage(rw http.ResponseWriter, r *http.Request) {

}

func (mh *MashStage) UpdateMashStage(rw http.ResponseWriter, r *http.Request) {
}

func (mh *MashStage) DeleteMashStage(rw http.ResponseWriter, r *http.Request) {
}
