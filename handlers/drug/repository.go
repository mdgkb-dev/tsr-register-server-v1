package drug

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) create(item *models.Drug) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll(diagnosisIds []uuid.UUID) ([]*models.Drug, error) {
	items := make([]*models.Drug, 0)
	q := r.db().NewSelect().Model(&items).
		Relation("DrugRegimens.DrugRegimenBlocks.DrugRegimenBlockItems", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("drug_regimen_block_items.order_item")
		}).
		Relation("DrugRegimens.DrugRegimenBlocks", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("drug_regimen_blocks.order_item")
		}).
		Relation("DrugsDiagnosis.MkbDiagnosis.MkbSubDiagnosis").
		Relation("DrugsDiagnosis.MkbDiagnosis.MkbGroup").
		Relation("DrugsDiagnosis.MkbSubDiagnosis").
		Relation("DrugsDiagnosis.MkbConcreteDiagnosis")

	if len(diagnosisIds) > 0 {
		q.Join("JOIN drugs_diagnosis ON drugs_diagnosis.drug_id = drugs.id").
			Where("drugs_diagnosis.mkb_diagnosis_id in (?)", bun.In(diagnosisIds))
	}
	err := q.Order("drugs.name").Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Drug, error) {
	item := models.Drug{}
	err := r.db().NewSelect().Model(&item).
		Relation("DrugRegimens.DrugRegimenBlocks.DrugRegimenBlockItems", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("drug_regimen_block_items.order_item")
		}).
		Relation("DrugRegimens.DrugRegimenBlocks", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("drug_regimen_blocks.order_item")
		}).
		Relation("DrugsDiagnosis.MkbDiagnosis.MkbSubDiagnosis").
		Relation("DrugsDiagnosis.MkbDiagnosis.MkbGroup").
		Relation("DrugsDiagnosis.MkbSubDiagnosis").
		Relation("DrugsDiagnosis.MkbConcreteDiagnosis").
		Where("id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Drug{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Drug) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
