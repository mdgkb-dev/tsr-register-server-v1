package models

import (
	"sort"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterPropertySet struct {
	bun.BaseModel                   `bun:"register_property_set,alias:register_property_set"`
	ID                              uuid.NullUUID          `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                            string                 `json:"name"`
	RegisterPropertyID              uuid.UUID              `bun:"type:uuid" json:"registerPropertyID"`
	Order                           int                    `bun:"register_property_set_order" json:"order"`
	RegisterPropertyOthers          RegisterPropertyOthers `bun:"rel:has-many" json:"registerPropertyOthers"`
	RegisterPropertyOthersForDelete []uuid.UUID            `bun:"-" json:"registerPropertyOthersForDelete"`

	AggregatedValues         map[string]float64       `bun:"-" json:"aggregatedValues"`
	RegisterQueryPercentages RegisterQueryPercentages `bun:"-" `
}

type RegisterPropertySets []*RegisterPropertySet

func (item *RegisterPropertySet) SetIDForChildren() {
	if len(item.RegisterPropertyOthers) == 0 {
		return
	}
	for i := range item.RegisterPropertyOthers {
		item.RegisterPropertyOthers[i].RegisterPropertySetID = item.ID
	}
}

func (items RegisterPropertySets) SetIDForChildren() {
	if len(items) == 0 {
		return
	}
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (items RegisterPropertySets) GetRegisterPropertyOthers() RegisterPropertyOthers {
	itemsForGet := make(RegisterPropertyOthers, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyOthers...)
	}
	return itemsForGet
}

func (items RegisterPropertySets) GetRegisterPropertyOthersForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyOthersForDelete...)
	}
	return itemsForGet
}

func (items RegisterPropertySets) Include(setToPatient []*RegisterPropertySetToPatient) string {
	exists := No
	for _, item := range items {
		if len(setToPatient) == 0 {
			continue
		}
		for _, s := range setToPatient {
			if s.RegisterPropertySetID == item.ID.UUID {
				exists = Yes
				break
			}
		}
	}
	return exists
}

func (item *RegisterPropertySet) writeXlsxAggregatedValues(key string) {
	_, ok := item.AggregatedValues[key]
	if ok {
		item.AggregatedValues[key]++
	} else {
		item.AggregatedValues[key] = 1
	}
}

func (item *RegisterPropertySet) GetAggregatedPercentage() {
	sum := float64(0)
	for k, v := range item.AggregatedValues {
		sum += v
		item.RegisterQueryPercentages = append(item.RegisterQueryPercentages, &RegisterQueryPercentage{k, v})
	}
	sort.Slice(item.RegisterQueryPercentages, func(i, j int) bool {
		return item.RegisterQueryPercentages[i].Value > item.RegisterQueryPercentages[j].Value
	})
}
