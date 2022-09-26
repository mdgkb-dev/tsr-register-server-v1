package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterProperty struct {
	bun.BaseModel                   `bun:"register_property,alias:register_property"`
	ID                              uuid.UUID              `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                            string                 `json:"name"`
	ShortName                       string                 `json:"shortName"`
	ColWidth                        int                    `json:"colWidth"`
	WithOther                       bool                   `json:"withOther"`
	Tag                             string                 `json:"tag"`
	Order                           int                    `bun:"register_property_order" json:"order"`
	ValueType                       *ValueType             `bun:"rel:belongs-to" json:"valueType"`
	ValueTypeID                     uuid.UUID              `bun:"type:uuid" json:"valueTypeId"`
	RegisterGroupID                 uuid.UUID              `bun:"type:uuid" json:"registerGroupId"`
	RegisterGroup                   *RegisterGroup         `bun:"rel:belongs-to" json:"registerGroup"`
	AgeCompare                      bool                   `json:"ageCompare"`
	RegisterPropertyRadios          RegisterPropertyRadios `bun:"rel:has-many" json:"registerPropertyRadios"`
	RegisterPropertyRadiosForDelete []uuid.UUID            `bun:"-" json:"registerPropertyRadiosForDelete"`

	RegisterPropertyVariants          RegisterPropertyVariants `bun:"rel:has-many" json:"registerPropertyVariants"`
	RegisterPropertyVariantsForDelete []uuid.UUID              `bun:"-" json:"registerPropertyVariantsForDelete"`

	RegisterPropertySets          RegisterPropertySets `bun:"rel:has-many" json:"registerPropertySets"`
	RegisterPropertySetsForDelete []uuid.UUID          `bun:"-" json:"registerPropertySetsForDelete"`

	RegisterPropertyExamples          RegisterPropertyExamples `bun:"rel:has-many" json:"registerPropertyExamples"`
	RegisterPropertyExamplesForDelete []uuid.UUID              `bun:"-" json:"registerPropertyExamplesForDelete"`

	RegisterPropertyMeasures          RegisterPropertyMeasures `bun:"rel:has-many" json:"registerPropertyMeasures"`
	RegisterPropertyMeasuresForDelete []uuid.UUID              `bun:"-" json:"registerPropertyMeasuresForDelete"`
}

type RegisterProperties []*RegisterProperty

func (item *RegisterProperty) SetIDForChildren() {
	if len(item.RegisterPropertyRadios) > 0 {
		for i := range item.RegisterPropertyRadios {
			item.RegisterPropertyRadios[i].RegisterPropertyID = item.ID
		}
	}
	if len(item.RegisterPropertySets) > 0 {
		for i := range item.RegisterPropertySets {
			item.RegisterPropertySets[i].RegisterPropertyID = item.ID
		}
	}
	if len(item.RegisterPropertyExamples) > 0 {
		for i := range item.RegisterPropertyExamples {
			item.RegisterPropertyExamples[i].RegisterPropertyID = item.ID
		}
	}
	if len(item.RegisterPropertyMeasures) > 0 {
		for i := range item.RegisterPropertyMeasures {
			item.RegisterPropertyMeasures[i].RegisterPropertyID = item.ID
		}
	}
	if len(item.RegisterPropertyVariants) > 0 {
		for i := range item.RegisterPropertyVariants {
			item.RegisterPropertyVariants[i].RegisterPropertyID = item.ID
		}
	}
}

func (items RegisterProperties) SetIDForChildren() {
	if len(items) == 0 {
		return
	}
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (items RegisterProperties) GetRegisterPropertyExamples() RegisterPropertyExamples {
	itemsForGet := make(RegisterPropertyExamples, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyExamples...)
	}
	return itemsForGet
}

func (items RegisterProperties) GetRegisterPropertyRadios() RegisterPropertyRadios {
	itemsForGet := make(RegisterPropertyRadios, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyRadios...)
	}
	return itemsForGet
}

func (items RegisterProperties) GetRegisterPropertySets() RegisterPropertySets {
	itemsForGet := make(RegisterPropertySets, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertySets...)
	}
	return itemsForGet
}

func (items RegisterProperties) GetRegisterPropertyRadioForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyRadiosForDelete...)
	}
	return itemsForGet
}

func (items RegisterProperties) GetRegisterPropertySetForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertySetsForDelete...)
	}
	return itemsForGet
}

func (items RegisterProperties) GetRegisterPropertyExamplesForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyExamplesForDelete...)
	}
	return itemsForGet
}

func (items RegisterProperties) GetRegisterPropertyMeasuresForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyMeasuresForDelete...)
	}
	return itemsForGet
}

func (items RegisterProperties) GetRegisterPropertyMeasures() RegisterPropertyMeasures {
	itemsForGet := make(RegisterPropertyMeasures, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyMeasures...)
	}
	return itemsForGet
}

func (items RegisterProperties) GetRegisterPropertyVariants() RegisterPropertyVariants {
	itemsForGet := make(RegisterPropertyVariants, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyVariants...)
	}
	return itemsForGet
}

func (items RegisterProperties) GetRegisterPropertyVariantsForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyVariantsForDelete...)
	}
	return itemsForGet
}
