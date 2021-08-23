package anthropometryData

import (
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) createMany(items []*models.AnthropometryData) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.AnthropometryData)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.AnthropometryData) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("anthropometry_id = EXCLUDED.anthropometry_id").
		Set("value = EXCLUDED.value").
		Set("date = EXCLUDED.date").
		Model(&items).
		Exec(r.ctx)
	return err
}
