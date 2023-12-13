package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Contact struct {
	bun.BaseModel `bun:"contact,alias:contact"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Phone         string        `json:"phone"`
	Email         string        `json:"email"`
}
type Contacts []*Contact

type ContactsWithCount struct {
	Contacts Contacts `json:"items"`
	Count    int      `json:"count"`
}
