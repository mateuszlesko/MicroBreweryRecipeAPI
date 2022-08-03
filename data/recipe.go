package data

import "time"

type Recipe struct {
	Id             int
	RecipeName     string
	CreatedAt      time.Time
	RecipeCategory *RecipeCategory
}
