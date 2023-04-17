package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MkbItem struct {
	bun.BaseModel `bun:"mkb_items,alias:mkb_items"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Code          string        `json:"code"`
	RangeStart    string        `json:"rangeStart"`
	RangeEnd      string        `json:"rangeEnd"`
	Comment       string        `json:"comment"`
	Children      MkbItems      `bun:"-" json:"children"`
	ParentID      uuid.NullUUID `bun:"type:uuid" json:"parentId"`
	Parent        *MkbItem      `bun:"-" json:"parent"`
	Leaf          bool          `json:"leaf"`
	Relevant      bool          `json:"relevant"`
}

type MkbItems []*MkbItem

func MkbListToTree(mkbItems MkbItems) *MkbItem {
	subtrees := map[uuid.NullUUID]*MkbItem{}
	for _, emp := range mkbItems {
		subtrees[emp.ParentID] = emp
	}
	for _, emp := range mkbItems {
		if emp.ParentID.Valid {
			subtree := subtrees[emp.ParentID]
			subtree.Children = append(subtree.Children, emp)
		}
	}
	return subtrees[uuid.NullUUID{}]
}

func (items MkbItems) GetItemsByParentID(parentID uuid.NullUUID) MkbItems {
	itemsForGet := make(MkbItems, 0)
	for i := range items {
		if items[i].ParentID == parentID {
			itemsForGet = append(itemsForGet, items[i])
		}
	}
	return itemsForGet
}

func (items MkbItems) GetRoot() *MkbItem {
	var item *MkbItem
	for i := range items {
		if !items[i].ParentID.Valid {
			item = items[i]
			break
		}
	}
	return item
}

func BuildTree(tree *MkbItem, items MkbItems) {
	itemsOnThisLevel := items.GetItemsByParentID(tree.ID)
	for i := range itemsOnThisLevel {
		tree.Children = append(tree.Children, itemsOnThisLevel[i])
		BuildTree(itemsOnThisLevel[i], items)
	}
}
