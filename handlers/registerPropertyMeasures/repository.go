package registerPropertyMeasures

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.RegisterPropertyMeasure) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items []*models.RegisterPropertyMeasure, err error) {
	err = r.db().NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.RegisterPropertyMeasure, error) {
	item := models.RegisterPropertyMeasure{}
	err := r.db().NewSelect().Model(&item).Where("register_property_measures.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.RegisterPropertyMeasure{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.RegisterPropertyMeasure) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.RegisterPropertyMeasures) (err error) {
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
		Model((*models.RegisterPropertyMeasure)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
