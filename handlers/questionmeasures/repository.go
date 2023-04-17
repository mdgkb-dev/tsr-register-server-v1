package questionmeasures

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.QuestionMeasure) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items []*models.QuestionMeasure, err error) {
	err = r.db().NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.QuestionMeasure, error) {
	item := models.QuestionMeasure{}
	err := r.db().NewSelect().Model(&item).Where("register_property_measures.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.QuestionMeasure{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.QuestionMeasure) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.QuestionMeasures) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set(`name = EXCLUDED.name`).
		Set(`register_property_id = EXCLUDED.register_property_id`).
		Set(`register_property_measure_order = EXCLUDED.register_property_measure_order`).
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.QuestionMeasure)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
