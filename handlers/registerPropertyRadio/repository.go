package registerPropertyRadio

import "mdgkb/tsr-tegister-server-v1/models"

func (r *Repository) create(item *models.RegisterPropertyRadio) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items []*models.RegisterPropertyRadio, err error) {
	err = r.db.NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.RegisterPropertyRadio, error) {
	item := models.RegisterPropertyRadio{}
	err := r.db.NewSelect().Model(&item).Where("RegisterPropertyRadio.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.RegisterPropertyRadio{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.RegisterPropertyRadio) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
