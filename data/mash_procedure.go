package data

type MashProcedureRecord struct {
	Temperature float32 `json:"T"`
	Holding     int32   `json:"H"`
}

type MashProcedure struct {
	RecipeId          int32                 `json:"RID"`
	ProcedureCount    int32                 `json:"PC"`
	MashProcedureList []MashProcedureRecord `json:"MPL"`
}
