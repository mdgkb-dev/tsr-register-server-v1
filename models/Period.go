package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Period struct {
	bun.BaseModel `bun:"period,alias:period"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	DateStart     time.Time `json:"dateStart"`
	DateEnd       time.Time `json:"dateEnd"`
}
