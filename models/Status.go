package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Status struct {
	bun.BaseModel `bun:"statuses,alias:statuses"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Color         string        `json:"color"`
	Model         string        `json:"model"`
}

type Statuses []*Status
type StatusesWithCount struct {
	Statuses Statuses `json:"items"`
	Count    int      `json:"count"`
}
