package patientDiagnosis

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items []*models.PatientDiagnosis) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.PatientDiagnosis)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.PatientDiagnosis) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("mkb_diagnosis_id = EXCLUDED.mkb_diagnosis_id").
		Set("mkb_sub_diagnosis_id = EXCLUDED.mkb_sub_diagnosis_id").
		Set(`"primary" = EXCLUDED."primary"`).
		Model(&items).
		Exec(r.ctx)
	return err
}
