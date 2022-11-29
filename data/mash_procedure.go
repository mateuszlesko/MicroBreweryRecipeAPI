package data

type MashProcedure struct {
	MashId               int32   `json:"MID"`
	ProcedureCount       int32   `json:"PC"`
	MashTemperaturesList []int32 `json:"MTpL"`
	MashTimeList         []int32 `json:"MTmL"`
}
