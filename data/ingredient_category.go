package data

import "time"

type IngredientCategory struct {
	Category_id         int       `json:"id"`
	Category_name       string    `json:"name"`
	Category_created_at time.Time `json:"createdAt"`
}
