package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CommissionStatus struct {
	bun.BaseModel `bun:"commissions_statuses,alias:commissions_statuses"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Color         string        `json:"color"`
}

type CommissionsStatuses []*CommissionStatus
type CommissionsStatusesWithCount struct {
	CommissionsStatuses CommissionsStatuses `json:"items"`
	Count               int                 `json:"count"`
}
