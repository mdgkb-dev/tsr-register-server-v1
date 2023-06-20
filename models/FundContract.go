package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FundContract struct {
	bun.BaseModel `bun:"fund_contracts,alias:fund_contracts"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Commissions   Commissions   `bun:"rel:has-many" json:"commissions"`
	DrugArrives   DrugArrives   `bun:"rel:has-many" json:"drugArrives"`
	Date          *time.Time    `bun:"item_date" json:"date"`
	Number        string        `json:"number"`

	DrugApplication   *DrugApplication `bun:"rel:belongs-to" json:"drugApplication"`
	DrugApplicationID uuid.NullUUID    `bun:"type:uuid" json:"drugApplicationId"`
}

type FundContracts []*FundContract
type FundContractsWithCount struct {
	FundContracts FundContracts `json:"items"`
	Count         int           `json:"count"`
}
