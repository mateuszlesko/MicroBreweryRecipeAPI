package data

type MashProcedureRecord struct {
	Temperature float32 `json:"T"`
	Holding     int32   `json:"H"`
}

type MashProcedure struct {
	MashId               int32   `json:"MID"`
	ProcedureCount       int32   `json:"PC"`
	MashTemperaturesList []int32 `json:"MTpL"`
	MashTimeList         []int32 `json:"MTmL"`
}
