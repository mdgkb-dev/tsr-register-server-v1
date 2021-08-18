package anthropometry

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) create(item *models.Anthropometry) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items []*models.Anthropometry, err error) {
	err = r.db.NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Anthropometry, error) {
	item := models.Anthropometry{}
	err := r.db.NewSelect().Model(&item).Where("anthropometry.id = ?", *id).Scan(context.Background())
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Anthropometry{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Anthropometry) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
