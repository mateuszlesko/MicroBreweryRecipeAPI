package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryRecipeAPI/data"
)

var mash_tum_ready bool = false
var mash_procedure_finish bool = false

type Mashing struct {
	l *log.Logger
}

func NewMashing(l *log.Logger) *Mashing {
	return &Mashing{l}
}

func (Mashing) GetProcedureToDo(rw http.ResponseWriter, r *http.Request) {
	mp := data.MashProcedure{
		MashId:               0x01,
		RecipeId:             1,
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

func (Mashing) GetMashTumReadiness(rw http.ResponseWriter, r *http.Request) {
	mash_tum_ready = true
	fmt.Println("%t", mash_tum_ready)
}

func (Mashing) GetMashingRaport(rw http.ResponseWriter, r *http.Request) {
	recipe_id := r.URL.Query().Get("rid")
	stage_id := r.URL.Query().Get("sid")
	sensor_bottom := r.URL.Query().Get("bs")
	sensor_top := r.URL.Query().Get("ts")
	control_signals := r.URL.Query().Get("cs")
	temperature_time_holding := r.URL.Query().Get("cl")
	fmt.Println("%d %d %d C %d C %d %d min\n", recipe_id, stage_id, sensor_bottom, sensor_top, control_signals, temperature_time_holding)
}

func (Mashing) GetRemoteControl(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(`{"cs":0}`)
}

func (Mashing) GetMashRemoteControl(rw http.ResponseWriter, r *http.Request) {
	r_c := data.RemoteControl{Control_signals: 0}
	r_c_bytes, err := json.Marshal(r_c)
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusUnprocessableEntity)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(r_c_bytes)
}

func (Mashing) GetMashProcedureFinish(rw http.ResponseWriter, r *http.Request) {
	mash_procedure_finish = true
	fmt.Println("done %t", mash_procedure_finish)
}
