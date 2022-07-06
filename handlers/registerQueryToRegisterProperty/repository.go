package registerQueryToRegisterProperty

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items []*models.RegisterQueryToRegisterProperty) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.RegisterQueryToRegisterProperty)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.RegisterQueryToRegisterProperty) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set(`"order" = EXCLUDED."order"`).
		Set(`register_query_id = EXCLUDED.register_query_id`).
		Set(`register_property_id = EXCLUDED.register_property_id`).
		Exec(r.ctx)
	return err
}
