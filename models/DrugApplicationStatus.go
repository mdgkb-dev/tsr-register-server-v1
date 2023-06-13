package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugApplicationStatus struct {
	bun.BaseModel `bun:"drug_applications_statuses,alias:drug_applications_statuses"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Color         string        `json:"color"`
}

type DrugApplicationsStatuses []*DrugApplicationStatus
type DrugApplicationsStatusesWithCount struct {
	DrugApplicationsStatuses DrugApplicationsStatuses `json:"items"`
	Count                    int                      `json:"count"`
}
