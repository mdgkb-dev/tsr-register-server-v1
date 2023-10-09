package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugRecipe struct {
	bun.BaseModel `bun:"drug_recipes,alias:drug_recipes"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Drug   *Drug         `bun:"rel:belongs-to" json:"drug"`
	DrugID uuid.NullUUID `bun:"type:uuid" json:"drugId"`

	DrugForm   *DrugForm     `bun:"rel:belongs-to" json:"drugForm"`
	DrugFormID uuid.NullUUID `bun:"type:uuid" json:"drugFormId"`

	DrugDoze   *DrugDoze     `bun:"rel:belongs-to" json:"drugDoze"`
	DrugDozeID uuid.NullUUID `bun:"type:uuid" json:"drugDozeId"`
}

type DrugRecipes []*DrugRecipe

type DrugRecipesWithCount struct {
	DrugRecipes DrugRecipes `json:"items"`
	Count       int         `json:"count"`
}
