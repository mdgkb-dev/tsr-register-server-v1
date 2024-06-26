package models

import (
	"time"

	"github.com/pro-assistance/pro-assister/helpers/uploader"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Disability struct {
	bun.BaseModel `bun:"disabilities,alias:disabilities"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Patient       *Patient      `bun:"rel:belongs-to" json:"patient"`
	PatientID     uuid.NullUUID `bun:"type:uuid" json:"patientId"`
	Edvs          Edvs          `bun:"rel:has-many" json:"edvs"`
	EdvsForDelete []uuid.UUID   `bun:"-" json:"edvsForDelete"`
	DateStart     time.Time     `json:"dateStart"`
	DateEnd       time.Time     `json:"dateEnd"`
	Series        string        `json:"series"`
	Number        string        `json:"number"`
	DeletedAt     *time.Time    `bun:",soft_delete" json:"deletedAt"`
}

type Disabilities []*Disability

type DisabilitiesWithCount struct {
	Disabilities Disabilities `json:"items"`
	Count        int          `json:"count"`
}

func (item *Disability) SetIDForChildren() {
	if len(item.Edvs) > 0 {
		for i := range item.Edvs {
			item.Edvs[i].DisabilityID = item.ID
		}
	}
}

func (item *Disability) SetFilePath(fileID *string) *string {
	for i := range item.Edvs {
		if item.Edvs[i].FileInfo.ID.UUID.String() == *fileID {
			item.Edvs[i].FileInfo.FileSystemPath = uploader.BuildPath(fileID)
			return &item.Edvs[i].FileInfo.FileSystemPath
		}
	}
	return nil
}

func GetEdvs(disabilities []*Disability) []*Edv {
	items := make([]*Edv, 0)
	if len(disabilities) == 0 {
		return items
	}
	for i := range disabilities {
		disabilities[i].SetIDForChildren()
		items = append(items, disabilities[i].Edvs...)
	}
	return items
}

func GetPeriodsFromDisability(items []*Disability) []*Period {
	itemsForGet := make([]*Period, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	//for i := range items {
	//	itemsForGet = append(itemsForGet, items[i].Period)
	//}
	return itemsForGet
}

func SetPeriodIDToDisabilities(items []*Disability) {
	if len(items) == 0 {
		return
	}
	//for i := range items {
	//	items[i].PeriodID = items[i].Period.ID
	//}
}
