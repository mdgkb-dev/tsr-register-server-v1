package patient

import (
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *HistoryRepository) getDB() *bun.DB {
	return r.db
}

func (r *HistoryRepository) create(item *models.PatientHistory) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *HistoryRepository) getAll(id *string) (items []*models.PatientHistory, err error) {
	err = r.db.NewSelect().
		Model(&items).
		Relation("History").
		Relation("HumanHistory").
		Relation("UpdatedBy").
		Where("patient_histories.id = ?", *id).WhereAllWithDeleted().Scan(r.ctx)
	return items, err
}

func (r *HistoryRepository) get(id *string) (*models.PatientHistory, error) {
	item := models.PatientHistory{}
	err := r.db.NewSelect().
		Model(&item).
		Relation("History").
		Relation("HumanHistory").
		Where("patient_histories.patient_history_id = ?", *id).WhereAllWithDeleted().Scan(r.ctx)
	return &item, err
}
