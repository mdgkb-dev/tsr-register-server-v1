package representativeToPatient

import (
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) createMany(items []*models.RepresentativeToPatient) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.RepresentativeToPatient)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.RepresentativeToPatient) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("representative_type_id = EXCLUDED.representative_type_id").
		Set("representative_id = EXCLUDED.representative_id").
		Model(&items).
		Exec(r.ctx)
	return err
}