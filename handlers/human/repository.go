package human

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Human) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Human) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("human.id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id uuid.UUID) (err error) {
	_, err = r.db.NewDelete().Model(&models.Human{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) get(id uuid.UUID) (*models.Human, error) {
	item := models.Human{}
	err := r.db.NewSelect().Model(&item).Where("human.id = ?", id).Scan(r.ctx)
	return &item, err
}
