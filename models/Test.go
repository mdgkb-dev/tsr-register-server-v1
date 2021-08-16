package models

import (
    "github.com/google/uuid"
)

type Test struct {
    ID         uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
    Name       string    `json:"name"`
}
