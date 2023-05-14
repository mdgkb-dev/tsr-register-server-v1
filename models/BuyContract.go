package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type BuyContract struct {
	bun.BaseModel  `bun:"buy_contracts,alias:buy_contracts"`
	ID             uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	FundContract   *FundContract `bun:"rel:belongs-to" json:"fundContract"`
	FundContractID uuid.NullUUID `bun:"type:uuid" json:"fundContractId"`
}

type BuyContracts []*BuyContract
type BuyContractsWithCount struct {
	BuyContracts BuyContracts `json:"items"`
	Count        int          `json:"count"`
}
