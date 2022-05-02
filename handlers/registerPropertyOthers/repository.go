package registerPropertyOthers

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.RegisterPropertyOther) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items []*models.RegisterPropertyOther, err error) {
	err = r.db.NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.RegisterPropertyOther, error) {
	item := models.RegisterPropertyOther{}
	err := r.db.NewSelect().Model(&item).Where("RegisterPropertyOther.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.RegisterPropertyOther{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.RegisterPropertyOther) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.RegisterPropertyOthers) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set(`name = EXCLUDED.name`).
		Set(`register_property_id = EXCLUDED.register_property_id`).
		Set(`register_property_radio_id = EXCLUDED.register_property_radio_id`).
		Set(`register_property_set_id = EXCLUDED.register_property_set_id`).
		Set(`register_property_others_order = EXCLUDED.register_property_others_order`).
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.RegisterPropertyOther)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
