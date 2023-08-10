package answers

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.Answer) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items []*models.Answer, err error) {
	err = r.db().NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Answer, error) {
	item := models.Answer{}
	err := r.db().NewSelect().Model(&item).Where("?TableAlias.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Answer{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Answer) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Answers) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("value_string = EXCLUDED.value_string").
		Set("value_number = EXCLUDED.value_number").
		Set("value_date = EXCLUDED.value_date").
		Set("value_other = EXCLUDED.value_other").
		Set("answer_variant_id = EXCLUDED.answer_variant_id").
		Set("question_id = EXCLUDED.question_id").
		Set("patient_id = EXCLUDED.patient_id").
		Set("research_result_id = EXCLUDED.research_result_id").
		Set("filled = EXCLUDED.filled").
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.Answer)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
