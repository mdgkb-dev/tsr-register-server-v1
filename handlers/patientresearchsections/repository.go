package patientresearchsections

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.ResearchResults) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.ResearchResult)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.ResearchResults) (err error) {
	_, err = r.db().NewInsert().
		On("conflict (id) do update").
		Set("register_groups_to_patients_date = EXCLUDED.register_groups_to_patients_date").
		Set("register_group_id = EXCLUDED.register_group_id").
		Set("patient_id = EXCLUDED.patient_id").
		Model(&items).
		Exec(r.ctx)
	return err
}
