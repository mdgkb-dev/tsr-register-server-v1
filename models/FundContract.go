package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FundContract struct {
	bun.BaseModel `bun:"fund_contracts,alias:fund_contracts"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Commissions   Commissions   `bun:"rel:has-many" json:"commissions"`
	BuyContracts  BuyContracts  `bun:"rel:has-many" json:"buyContracts"`
}

type FundContracts []*FundContract
type FundContractsWithCount struct {
	FundContracts FundContracts `json:"items"`
	Count         int           `json:"count"`
}
