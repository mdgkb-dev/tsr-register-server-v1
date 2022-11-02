package registerpropertytopatient

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.RegisterPropertiesToPatients) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.RegisterPropertyToPatient)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.RegisterPropertiesToPatients) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("value_string = EXCLUDED.value_string").
		Set("value_number = EXCLUDED.value_number").
		Set("value_date = EXCLUDED.value_date").
		Set("value_other = EXCLUDED.value_other").
		Set("register_property_radio_id = EXCLUDED.register_property_radio_id").
		Set("register_property_id = EXCLUDED.register_property_id").
		Set("register_property_measure_id = EXCLUDED.register_property_measure_id").
		Set("register_property_variant_id = EXCLUDED.register_property_variant_id").
		Model(&items).
		Exec(r.ctx)
	return err
}
