package registerPropertyToPatient

import (
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items []*models.RegisterPropertyToPatient) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.RegisterPropertyToPatient)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.RegisterPropertyToPatient) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("value_string = EXCLUDED.value_string").
		Set("value_number = EXCLUDED.value_number").
		Set("value_date = EXCLUDED.value_date").
		Set("value_other = EXCLUDED.value_other").
		Set("register_property_radio_id = EXCLUDED.register_property_radio_id").
		Set("register_property_id = EXCLUDED.register_property_id").
		Model(&items).
		Exec(r.ctx)
	return err
}
