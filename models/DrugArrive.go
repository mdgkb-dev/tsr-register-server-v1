package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugArrive struct {
	bun.BaseModel  `bun:"drug_arrives,alias:drug_arrives"`
	ID             uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	FundContract   *FundContract `bun:"rel:belongs-to" json:"fundContract"`
	FundContractID uuid.NullUUID `bun:"type:uuid" json:"fundContractId"`

	//Drug   *Drug         `bun:"rel:belongs-to" json:"drug"`
	//DrugID uuid.NullUUID `bun:"type:uuid" json:"drugId"`

	Quantity int  `json:"quantity"`
	Stage    int  `json:"stage"`
	Arrived  bool `json:"arrived"`

	Date *time.Time `bun:"item_date" json:"date"`

	DrugDecreases DrugDecreases `bun:"rel:has-many" json:"drugDecreases"`
}

type DrugArrives []*DrugArrive
type DrugArrivesWithCount struct {
	DrugArrives DrugArrives `json:"items"`
	Count       int         `json:"count"`
}
