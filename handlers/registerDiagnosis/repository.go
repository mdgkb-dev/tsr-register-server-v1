package registerDiagnosis

import (
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) createMany(items []*models.RegisterDiagnosis) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.RegisterDiagnosis)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.RegisterDiagnosis) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set(`mkb_diagnosis_id = EXCLUDED.mkb_diagnosis_id`).
		Set(`mkb_sub_diagnosis_id = EXCLUDED.mkb_sub_diagnosis_id`).
		Exec(r.ctx)
	return err
}
