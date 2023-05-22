package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FundCouncil struct {
	bun.BaseModel `bun:"fund_councils,alias:fund_councils"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Commissions   Commissions   `bun:"rel:has-many" json:"commissions"`
	Date          *time.Time    `bun:"item_date" json:"date"`
	Number        int           `json:"number"`
}

type FundCouncils []*FundCouncil
type FundCouncilsWithCount struct {
	FundCouncils FundCouncils `json:"items"`
	Count        int          `json:"count"`
}
