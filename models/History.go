package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type History struct {
	bun.BaseModel `bun:"histories,alias:histories"`
	ID            uuid.UUID    `bun:"type:uuid,nullzero,notnull,default:uuid_generate_v4()" json:"id" `
	RequestType   *RequestType `json:"requestType"`
	RequestDate   time.Time    `bun:",nullzero,notnull,default:current_timestamp" json:"requestDate"`
}

type RequestType string

const (
	RequestTypeInsert RequestType = "Создание"
	RequestTypeUpdate RequestType = "Обновление"
	RequestTypeDelete RequestType = "Удаление"
)
