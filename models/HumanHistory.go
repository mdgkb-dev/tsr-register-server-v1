package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type HumanHistory struct {
	bun.BaseModel  `bun:"human_histories,alias:human_histories"`
	HumanHistoryID uuid.UUID `bun:"type:uuid,nullzero,notnull,default:uuid_generate_v4()" json:"humanHistoryId" `
	Human
}
