package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Doctor struct {
	bun.BaseModel `bun:"doctors,alias:doctors"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Position      string        `json:"position"`
}

type Doctors []*Doctor
type DoctorsWithCount struct {
	Doctors Doctors `json:"items"`
	Count   int     `json:"count"`
}
