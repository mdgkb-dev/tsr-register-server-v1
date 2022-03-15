package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"sort"
)

type RegisterQuery struct {
	bun.BaseModel                            `bun:"register_queries,alias:register_queries"`
	ID                                       uuid.UUID                          `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Name                                     string                             `json:"name"`
	Type                                     string                             `bun:"type:register_query_type_enum" json:"type"`
	Register                                 *Register                          `bun:"rel:belongs-to" json:"register"`
	RegisterID                               uuid.UUID                          `bun:"type:uuid" json:"registerId"`
	RegisterQueryToRegisterProperty          []*RegisterQueryToRegisterProperty `bun:"rel:has-many" json:"registerQueryToRegisterProperty"`
	RegisterQueryToRegisterPropertyForDelete []string                           `bun:"-" json:"registerQueryToRegisterPropertyForDelete"`
	Data                                     []map[string]interface{}           `bun:"-"'`
	Keys                                     []string                           `bun:"-" json:"keys"`
	Key                                      string                             `json:"key"`
}

type RegisterQueries []*RegisterQuery

func (item *RegisterQuery) SortKeys() {
	fmt.Println(item.Data)
	if len(item.Data) == 0 {
		return
	}
	item.Keys = make([]string, 0, len(item.Data[0]))
	for k := range item.Data[0] {
		if k != item.Key {
			item.Keys = append(item.Keys, k)
		}
	}
	sort.Strings(item.Keys)
	item.Keys = append([]string{item.Key}, item.Keys...)
	fmt.Println(item.Data[0])
}

func (item *RegisterQuery) SetIdForChildren() {
	if len(item.RegisterQueryToRegisterProperty) == 0 {
		return
	}

	for i := range item.RegisterQueryToRegisterProperty {
		item.RegisterQueryToRegisterProperty[i].RegisterQueryID = item.ID
	}
}
