package models

import (
	"fmt"
	"mdgkb/tsr-tegister-server-v1/helpers/xlsxhelper"
	"sort"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterQueryGroup struct {
	bun.BaseModel   `bun:"register_query_groups,alias:register_query_groups"`
	ID              uuid.UUID      `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name            string         `bun:"-" json:"name"`
	RegisterQueryID uuid.UUID      `bun:"type:uuid" json:"registerQueryId"`
	RegisterQuery   *RegisterQuery `bun:"rel:belongs-to" json:"registerQuery"`

	RegisterGroupID uuid.UUID `bun:"type:uuid" json:"registerGroupId"`
	RegisterGroup   *Research `bun:"rel:belongs-to" json:"registerGroup"`

	RegisterQueryGroupProperties RegisterQueryGroupProperties `bun:"rel:has-many" json:"registerQueryGroupProperties"`

	Order int `bun:"item_order" json:"order"`

	AggregateType        AggregateType `json:"aggregateType"`
	AggregatedProperties Questions     `bun:"rel:has-many" json:"aggregatedProperties"`

	CountSum        bool `json:"countSum"`
	CountPercentage bool `json:"countPercentage"`

	Sum                      uint                     `bun:"-" json:"sum"`
	Percentage               uint                     `bun:"-" json:"percentage"`
	AggregatedValues         map[string]float64       `bun:"-" json:"aggregatedValues"`
	RegisterQueryPercentages RegisterQueryPercentages `bun:"-" `

	PatientIndex int `bun:"-"`
}

type RegisterQueryPercentage struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

type RegisterQueryPercentages []*RegisterQueryPercentage

type RegisterQueryGroups []*RegisterQueryGroup

func (item *RegisterQueryGroup) GetResultFromData(prop *Question, propertyIndex int) string {
	return item.getAggregatedData(prop, propertyIndex)
}

func (item *RegisterQueryGroup) getAggregatedData(prop *Question, propertyIndex int) string {
	//if item.AggregateType == AggregateNone {
	//	if prop.ValueType.IsSet() {
	//		return prop.RegisterPropertySets.Include(item.RegisterGroup.RegisterGroupsToPatients[item.PatientIndex].RegisterPropertySetToPatient)
	//	}
	//	if len(item.RegisterGroup.RegisterGroupsToPatients[item.PatientIndex].RegisterPropertyToPatient) > 0 {
	//		if propertyIndex < len(item.RegisterGroup.RegisterGroupsToPatients[item.PatientIndex].RegisterPropertyToPatient) {
	//			return item.RegisterGroup.RegisterGroupsToPatients[item.PatientIndex].RegisterPropertyToPatient[propertyIndex].GetData(prop)
	//		}
	//		return No
	//	}
	//}
	//if item.AggregateType == AggregateExisting {
	//	if item.RegisterGroup.RegisterGroupsToPatients != nil && item.PatientIndex < len(item.RegisterGroup.RegisterGroupsToPatients) {
	//		return item.RegisterGroup.RegisterGroupsToPatients[item.PatientIndex].GetAggregateExistingData()
	//	}
	//}
	return ""
}

func (item *RegisterQueryGroup) GetAggregatedPercentage() {
	sum := float64(0)
	for k, v := range item.AggregatedValues {
		sum += v
		item.RegisterQueryPercentages = append(item.RegisterQueryPercentages, &RegisterQueryPercentage{k, v})
	}
	sort.Slice(item.RegisterQueryPercentages, func(i, j int) bool {
		return item.RegisterQueryPercentages[i].Value > item.RegisterQueryPercentages[j].Value
	})
}

func (items RegisterQueryGroups) writeXlsxHeader(xl *xlsxhelper.XlsxHelper) {
	for i := range items {
		items[i].writeXlsxHeader(xl)
	}
}

func (item *RegisterQueryGroup) writeXlsxHeader(xl *xlsxhelper.XlsxHelper) {
	xl.WriteString(1, xl.Cursor, &[]string{item.Name})
	if item.AggregateType == AggregateNone {
		item.RegisterQueryGroupProperties.writeXlsxHeader(xl)
	} else {
		xl.WriteString(1, xl.Cursor, &[]string{item.Name})
		xl.Cursor++
	}
}

func (items RegisterQueryGroups) writeXlsxData(xl *xlsxhelper.XlsxHelper, id uuid.UUID) {
	for i := range items {
		writeEmpty := false
		//if items[i].RegisterGroup.RegisterGroupsToPatients[items[i].PatientIndex].PatientID != id {
		//	writeEmpty = true
		//}
		items[i].writeXlsxData(xl, writeEmpty)
		if !writeEmpty {
			items[i].PatientIndex++
		}
	}
}
func (item *RegisterQueryGroup) writeXlsxData(xl *xlsxhelper.XlsxHelper, writeEmpty bool) {
	if item.AggregateType == AggregateNone {
		item.RegisterQueryGroupProperties.writeXlsxData(xl, item, writeEmpty)
	}
	if item.AggregateType == AggregateExisting {
		res := item.GetResultFromData(nil, 0)
		if writeEmpty {
			res = NoData
		}
		str := fmt.Sprintf("%v", res)
		item.writePercentage(str)
		xl.Data = append(xl.Data, str)
	}
}

func (item *RegisterQueryGroup) writePercentage(key string) {
	_, ok := item.AggregatedValues[key]
	if ok {
		item.AggregatedValues[key]++
	} else {
		item.AggregatedValues[key] = 1
	}
}

func (items RegisterQueryGroups) writeAggregates(xl *xlsxhelper.XlsxHelper) {
	for i := range items {
		items[i].writeAggregates(xl)
	}
}

func (item *RegisterQueryGroup) writeAggregates(xl *xlsxhelper.XlsxHelper) {
	if item.AggregateType == AggregateNone {
		item.RegisterQueryGroupProperties.writeAggregates(xl)
	}
	if item.AggregateType == AggregateExisting {
		item.AggregateType.WriteAggregatedValues(xl, item.AggregatedValues)
	}
}
