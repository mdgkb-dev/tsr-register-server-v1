package contact

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) create(item *models.Contact) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Contact) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.Contact) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("phone = EXCLUDED.phone").
		Set("email = EXCLUDED.email").
		Model(item).
		Exec(r.ctx)
	return err
}

func (r *Repository) createMany(items []*models.Contact) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.Contact) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("phone = EXCLUDED.phone").
		Set("email = EXCLUDED.email").
		Model(&items).
		Exec(r.ctx)
	return err
}
