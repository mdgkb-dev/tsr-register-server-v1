package representative

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) create(item *models.Representative) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll(offset *int) (items []*models.Representative, err error) {
	err = r.db.NewSelect().
		Model(&items).
		Relation("Human").
		Offset(*offset * 25).
		Limit(25).
		Order("human.surname").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Representative, error) {
	item := models.Representative{}
	err := r.db.NewSelect().Model(&item).Where("id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Representative{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Representative) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
