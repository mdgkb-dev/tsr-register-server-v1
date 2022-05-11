package drug

import (
	"github.com/google/uuid"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Drug) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll(diagnosisIds []uuid.UUID) ([]*models.Drug, error) {
	items := make([]*models.Drug, 0)
	q := r.db.NewSelect().Model(&items).
		Relation("DrugRegimens.DrugRegimenBlocks.DrugRegimenBlockItems", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("drug_regimen_block_items.order_item")
		}).
		Relation("DrugRegimens.DrugRegimenBlocks", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("drug_regimen_blocks.order_item")
		}).
		Relation("DrugsDiagnosis.MkbDiagnosis.MkbSubDiagnosis").
		Relation("DrugsDiagnosis.MkbDiagnosis.MkbGroup").
		Relation("DrugsDiagnosis.MkbSubDiagnosis")

	if len(diagnosisIds) > 0 {
		q.Join("JOIN drugs_diagnosis ON drugs_diagnosis.drug_id = drugs.id").
			Where("drugs_diagnosis.mkb_diagnosis_id in (?)", bun.In(diagnosisIds))
	}
	err := q.Order("drugs.name").Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Drug, error) {
	item := models.Drug{}
	err := r.db.NewSelect().Model(&item).
		Relation("DrugRegimens.DrugRegimenBlocks.DrugRegimenBlockItems", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("drug_regimen_block_items.order_item")
		}).
		Relation("DrugRegimens.DrugRegimenBlocks", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("drug_regimen_blocks.order_item")
		}).
		Relation("DrugsDiagnosis.MkbDiagnosis.MkbSubDiagnosis").
		Relation("DrugsDiagnosis.MkbDiagnosis.MkbGroup").
		Relation("DrugsDiagnosis.MkbSubDiagnosis").
		Where("id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Drug{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Drug) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
