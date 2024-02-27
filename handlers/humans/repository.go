package humans

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) Create(c context.Context, item *models.Human) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll(c context.Context) (item models.HumansWithCount, err error) {
	item.Humans = make(models.Humans, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&item.Humans)

	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) Get(c context.Context, slug string) (*models.Human, error) {
	item := models.Human{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Where("?TableAlias.id = ?", slug).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Human{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Human) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
