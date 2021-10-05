package history

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) get(id *string) (*models.History, error) {
	item := models.History{}
	err := r.db.NewSelect().Model(&item).
		Where("history.id = ?", *id).Scan(r.ctx)

	return &item, err
}

func (r *Repository) create(item *models.History) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}
