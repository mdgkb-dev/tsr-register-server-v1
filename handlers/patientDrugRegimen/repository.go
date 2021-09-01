package patientDrugRegimen

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items []*models.PatientDrugRegimen) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.PatientDrugRegimen)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.PatientDrugRegimen) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("patient_id = EXCLUDED.patient_id").
		Model(&items).
		Exec(r.ctx)
	return err
}