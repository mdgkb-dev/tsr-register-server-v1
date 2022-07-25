package period

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.Period) (err error) {
	_, err = r.db().NewInsert().Model(&item).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Period) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) createMany(items []*models.Period) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.Period) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("date_start = EXCLUDED.date_start").
		Set("date_end = EXCLUDED.date_end").
		Model(&items).
		Exec(r.ctx)
	return err
}
