package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pro-assistance/pro-assister/uploadHelper"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Human struct {
	bun.BaseModel       `bun:"humans,alias:humans"`
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
	PhotoID             uuid.NullUUID `bun:"type:uuid" json:"photoId"`
	DeletedAt           *time.Time    `bun:",soft_delete" json:"deletedAt"`

	Documents          []*Document `bun:"rel:has-many" json:"documents"`
	DocumentsForDelete []uuid.UUID `bun:"-" json:"documentsForDelete"`

	InsuranceCompanyToHuman          []*InsuranceCompanyToHuman `bun:"rel:has-many" json:"insuranceCompanyToHuman"`
	InsuranceCompanyToHumanForDelete []uuid.UUID                `bun:"-" json:"insuranceCompanyToHumanForDelete"`
}

type Humans []*Human

type HumansWithCount struct {
	Humans Humans `json:"items"`
	Count  int    `json:"count"`
}

func (item *Human) SetFilePath(fileID *string) *string {
	for i := range item.Documents {
		path := item.Documents[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	if item.Photo != nil && item.Photo.ID.String() == *fileID {
		item.Photo.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.Photo.FileSystemPath
	}
	return nil
}

func (item *Human) SetIDForChildren() {
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

func (item *Human) SetDeleteIDForChildren() {
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

func (item *Human) GetFullName() string {
	return fmt.Sprintf("%s %s %s", item.Surname, item.Name, item.Patronymic)
}

func (item *Human) GetAge() int {
	return item.ageFromDateOfBirth()
}

//func (item *Human) GetFormattedAge() int {
//	return int(dur.Nanoseconds())
//}

func monthToInt(month string) int {
	switch month {
	case "January":
		return 1
	case "February":
		return 2
	case "March":
		return 3
	case "April":
		return 4
	case "May":
		return 5
	case "June":
		return 6
	case "July":
		return 7
	case "August":
		return 8
	case "September":
		return 9
	case "October":
		return 10
	case "November":
		return 11
	case "December":
		return 12
	default:
		panic("Unrecognized month")
	}
}

func (item *Human) ageFromDateOfBirth() int {
	var ageYear int
	ageYear = time.Now().Year() - item.DateBirth.Year()

	dobDayMonth, _ := strconv.Atoi(strconv.Itoa(item.DateBirth.Day()) + strconv.Itoa(monthToInt(item.DateBirth.Month().String())))
	nowDayMonth, _ := strconv.Atoi(strconv.Itoa(time.Now().Day()) + strconv.Itoa(monthToInt(time.Now().Month().String())))

	if dobDayMonth > nowDayMonth {
		ageYear = ageYear - 1
	}
	return ageYear
}

func (item *Human) GetFormattedDateBirth() string {
	return item.DateBirth.Format("02.01.2006")
}
