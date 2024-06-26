package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helpers/token"
	"github.com/pro-assistance/pro-assister/middleware"
)

type ModelInfo struct {
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updatedAt"`

	CreatedByID uuid.NullUUID `json:"createdById"`
	UpdatedByID uuid.NullUUID `json:"updatedById"`

	CreatedBy *User      `bun:"rel:belongs-to" json:"createdBy"`
	UpdatedBy *User      `bun:"rel:belongs-to" json:"updatedBy"`
	DeletedAt *time.Time `bun:",soft_delete" json:"deletedAt"`
}

func (item *ModelInfo) FillModelInfoUpdate(_ *gin.Context, _ *token.Token) error {
	//userID, err := tokenHelper.GetUserID(c)
	//if err != nil {
	//	return err
	//}
	//m.UpdatedByID = *userID
	//m.UpdatedAt = time.Now()
	return nil
}

func (item *ModelInfo) FillModelInfoCreate(c *gin.Context, tokenHelper *token.Token) (err error) {
	uid, err := tokenHelper.ExtractTokenMetadata(c.Request, middleware.ClaimUserID)
	if err != nil {
		return err
	}
	item.CreatedByID.UUID, err = uuid.Parse(uid)
	if err != nil {
		return err
	}
	item.CreatedByID.Valid = true
	if err != nil {
		return err
	}
	return nil
}
