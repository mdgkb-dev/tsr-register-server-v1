package drugRegimen

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items []*models.DrugRegimen) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.DrugRegimen)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.DrugRegimen) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("drug_id = EXCLUDED.drug_id").
		Set("name = EXCLUDED.name").
		Model(&items).
		Exec(r.ctx)
	return err
}
