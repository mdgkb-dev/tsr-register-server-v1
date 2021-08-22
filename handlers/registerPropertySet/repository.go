package registerPropertySet

import "mdgkb/tsr-tegister-server-v1/models"

func (r *Repository) create(item *models.RegisterPropertySet) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items []*models.RegisterPropertySet, err error) {
	err = r.db.NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.RegisterPropertySet, error) {
	item := models.RegisterPropertySet{}
	err := r.db.NewSelect().Model(&item).Where("RegisterPropertySet.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.RegisterPropertySet{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.RegisterPropertySet) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
