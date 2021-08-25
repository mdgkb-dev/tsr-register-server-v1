package registerGroupToRegister

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) createMany(items []*models.RegisterGroupToRegister) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.RegisterGroupToRegister)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.RegisterGroupToRegister) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set(`"order" = EXCLUDED."order"`).
		Set(`register_group_id = EXCLUDED.register_group_id`).
		Exec(r.ctx)
	return err
}
