package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/data"
)

type MashHistory struct {
	l *log.Logger
}

func NewMashHistory(l *log.Logger) *MashHistory {
	return &MashHistory{l}
}

func (MashHistory) GetAllMashings(rw http.ResponseWriter, r *http.Request) {
	allMashings, err := data.SelectMashHistory()
	if err != nil {
		http.Error(rw, "unable to get data", http.StatusBadRequest)
	}
	mashingsBytes, err := json.MarshalIndent(allMashings, "", "\t")
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(mashingsBytes)
}

func (MashHistory) GetCurrentMashings(rw http.ResponseWriter, r *http.Request) {
	currentMashs, err := data.SelectCurrentMashings()
	if err != nil {
		http.Error(rw, "unable to get data", http.StatusBadRequest)
	}
	mashsBytes, err := json.MarshalIndent(currentMashs, "", "\t")
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(mashsBytes)
}

func PostMashHistory(rw http.ResponseWriter, r *http.Request) {
	recipeId, err := strconv.Atoi(r.URL.Query().Get("recipeId"))
	if err != nil || recipeId == 0 {
		log.Panic()
		http.Error(rw, "unable to parse value", http.StatusBadRequest)
	}
	latestMashingId, err := data.InsertMashHistory(recipeId)
	if err != nil && latestMashingId < 0 {
		http.Error(rw, "Unable to add record to mash history", http.StatusBadRequest)
	}
	binaryLatestId := byte(latestMashingId)
	rw.Write([]byte{binaryLatestId})
}

func (MashHistory) PatchMashHistory(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id == 0 {
		log.Panic()
		http.Error(rw, "unable to parse value", http.StatusBadRequest)
	}
	err = data.UpdateMashHistory(id)
	if err != nil {
		http.Error(rw, "Unable to add record to mash history", http.StatusBadRequest)
	}
}
