package data

import "time"

type MashStage struct {
	Id          int
	StageTime   int64
	Temperature float32
	CreatedAt   time.Time
	Recipe      *Recipe
}
