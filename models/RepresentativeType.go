package models

import (
	"github.com/google/uuid"
)

type RepresentativeType struct {
	ID             *uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name           *string    `json:"name"`
	ChildMaleType  *string    `json:"childMaleType"`
	ChildWomanType *string    `json:"childWomanType"`
}
