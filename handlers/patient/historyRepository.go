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
	err = r.db.NewSelect().Model(&items).
	Relation("History").
	Relation("Human", func (q *bun.SelectQuery) *bun.SelectQuery {
		return q.ModelTableExpr("human_histories")
	}).
	Where("patient_histories.id = ?", *id).Scan(r.ctx)
	return items, err
}
