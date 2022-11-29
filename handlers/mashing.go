package handlers

import (
	"bytes"
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
		MashId:               0x01,
		ProcedureCount:       2,
		MashTemperaturesList: []int32{50, 55},
		MashTimeList:         []int32{15, 12},
	}
	mashingBytes, err := json.MarshalIndent(mp, "", "")
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
	}
	dst := &bytes.Buffer{}
	if json.Compact(dst, mashingBytes); err != nil {
		http.Error(rw, "unable to compress", http.StatusUnprocessableEntity)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(dst.Bytes())
}
