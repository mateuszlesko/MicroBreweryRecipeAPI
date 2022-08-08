package data

import (
	"encoding/json"
	"io"
	"time"
)

type IngredientCategory struct {
	Category_id         int       `json:"id"`
	Category_name       string    `json:"name"`
	Category_created_at time.Time `json:"createdAt"`
}

type IngredientCategoryFormVM struct {
	Category_id   int    `json:"id"`
	Category_name string `json:"name"`
}

func (c *IngredientCategory) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func (c *IngredientCategory) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(c) // pass reference to myself, map json to struct
}
