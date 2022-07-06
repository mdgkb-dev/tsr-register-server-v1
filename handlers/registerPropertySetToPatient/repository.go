package registerPropertySetToPatient

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items []*models.RegisterPropertySetToPatient) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.RegisterPropertySetToPatient)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.RegisterPropertySetToPatient) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("register_property_set_id = EXCLUDED.register_property_set_id").
		Set("prop_with_date_id = EXCLUDED.prop_with_date_id").
		Model(&items).
		Exec(r.ctx)
	return err
}
