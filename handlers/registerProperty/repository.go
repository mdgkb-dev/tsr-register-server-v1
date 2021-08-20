package registerProperty

import "mdgkb/tsr-tegister-server-v1/models"

func (r *Repository) create(item *models.RegisterProperty) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() ([]*models.RegisterProperty, error) {
	items := []*models.RegisterProperty{}
	err := r.db.NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.RegisterProperty, error) {
	item := models.RegisterProperty{}
	err := r.db.NewSelect().Model(&item).
		Relation("RegisterPropertyRadio").
		Relation("RegisterPropertySet").
		Where("register_property.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.RegisterProperty{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.RegisterProperty) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) getValueTypes() ([]*models.ValueType, error) {
	items := []*models.ValueType{}
	err := r.db.NewSelect().
		Model(&items).
		Scan(r.ctx)
	return items, err
}
