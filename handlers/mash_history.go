package handlers

import (
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

func PutMashHistory(rw http.ResponseWriter, r *http.Request) {

}
