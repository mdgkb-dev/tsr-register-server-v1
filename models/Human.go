package models

import (
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Human struct {
	bun.BaseModel       `bun:"human,select:humans_view,alias:human"`
	ID                  uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                string        `json:"name"`
	Surname             string        `json:"surname"`
	Patronymic          string        `json:"patronymic"`
	IsMale              bool          `json:"isMale"`
	DateBirth           *time.Time    `json:"dateBirth"`
	AddressRegistration string        `json:"addressRegistration"`
	AddressResidential  string        `json:"addressResidential"`
	Contact             *Contact      `bun:"rel:belongs-to" json:"contact"`
	ContactID           uuid.NullUUID `bun:"type:uuid" json:"contactId"`
	Photo               *FileInfo     `bun:"rel:belongs-to" json:"photo"`
	PhotoId             uuid.NullUUID `bun:"type:uuid" json:"photoId"`
	DeletedAt           *time.Time    `bun:",soft_delete" json:"deletedAt"`

	Documents          []*Document `bun:"rel:has-many" json:"documents"`
	DocumentsForDelete []uuid.UUID `bun:"-" json:"documentsForDelete"`

	InsuranceCompanyToHuman          []*InsuranceCompanyToHuman `bun:"rel:has-many" json:"insuranceCompanyToHuman"`
	InsuranceCompanyToHumanForDelete []uuid.UUID                `bun:"-" json:"insuranceCompanyToHumanForDelete"`
}

func (item *Human) SetFilePath(fileId *string) *string {
	for i := range item.Documents {
		path := item.Documents[i].SetFilePath(fileId)
		if path != nil {
			return path
		}
	}
	if item.Photo != nil && item.Photo.ID.String() == *fileId {
		item.Photo.FileSystemPath = uploadHelper.BuildPath(fileId)
		return &item.Photo.FileSystemPath
	}
	return nil
}

func (item *Human) SetIdForChildren() {
	if len(item.Documents) > 0 {
		for i := range item.Documents {
			item.Documents[i].HumanID = item.ID
		}
	}
	if len(item.InsuranceCompanyToHuman) > 0 {
		for i := range item.InsuranceCompanyToHuman {
			item.InsuranceCompanyToHuman[i].HumanID = item.ID
		}
	}
}

func (item *Human) SetDeleteIdForChildren() {
	for i := range item.Documents {
		item.DocumentsForDelete = append(item.DocumentsForDelete, item.Documents[i].ID)
	}
	for i := range item.InsuranceCompanyToHuman {
		item.InsuranceCompanyToHumanForDelete = append(item.InsuranceCompanyToHumanForDelete, item.InsuranceCompanyToHuman[i].ID)
	}
}

type InsuranceCompanyToHuman struct {
	bun.BaseModel      `bun:"insurance_company_to_human,alias:insurance_company_to_human"`
	ID                 uuid.UUID         `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Number             string            `json:"number"`
	InsuranceCompany   *InsuranceCompany `bun:"rel:belongs-to" json:"insuranceCompany"`
	InsuranceCompanyID uuid.UUID         `bun:"type:uuid" json:"insuranceCompanyId"`
	Human              *Human            `bun:"rel:belongs-to" json:"human"`
	HumanID            uuid.UUID         `bun:"type:uuid" json:"humanId"`
	DeletedAt          *time.Time        `bun:",soft_delete" json:"deletedAt"`
}
