package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Period struct {
	bun.BaseModel `bun:"period,alias:period"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	DateStart     time.Time `json:"dateStart"`
	DateEnd       time.Time `json:"dateEnd"`
}
