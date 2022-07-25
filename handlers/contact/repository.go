package contact

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"

	"github.com/google/uuid"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.Contact) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Contact) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.Contact) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("phone = EXCLUDED.phone").
		Set("email = EXCLUDED.email").
		Model(item).
		Exec(r.ctx)
	return err
}

func (r *Repository) delete(id uuid.NullUUID) (err error) {
	_, err = r.db().NewDelete().Model(&models.Contact{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) createMany(items []*models.Contact) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.Contact) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("phone = EXCLUDED.phone").
		Set("email = EXCLUDED.email").
		Model(&items).
		Exec(r.ctx)
	return err
}
