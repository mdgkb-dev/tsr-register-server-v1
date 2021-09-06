package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

type ModelInfo struct {
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updatedAt"`

	CreatedByID uuid.UUID `json:"createdById"`
	UpdatedByID uuid.UUID `json:"updatedById"`

	CreatedBy *User `bun:"rel:belongs-to" json:"createdBy"`
	UpdatedBy *User `bun:"rel:belongs-to" json:"updatedBy"`
}

func (m *ModelInfo) FillModelInfoUpdate(c *gin.Context) error {
	userId, err := GetUserID(c)
	if err != nil {
		return err
	}
	fmt.Println(userId)
	m.UpdatedByID = *userId
	return nil
}

func (m *ModelInfo) FillModelInfoCreate(c *gin.Context) error {
	userId, err := GetUserID(c)
	if err != nil {
		return err
	}
	m.CreatedByID = *userId
	m.UpdatedByID = *userId
	return nil
}
