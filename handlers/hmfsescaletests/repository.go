package hmfsescaletests

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.HmfseScaleTests) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.HmfseScaleTest)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.HmfseScaleTests) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("item_date = EXCLUDED.item_date").
		Set("patient_id = EXCLUDED.patient_id").
		Model(&items).
		Exec(r.ctx)
	return err
}
