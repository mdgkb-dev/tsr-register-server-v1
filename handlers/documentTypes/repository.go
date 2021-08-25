package documentTypes

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.DocumentType) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items []*models.DocumentType, err error) {
	err = r.db.NewSelect().Model(&items).Relation("DocumentTypeFields").Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.DocumentType, error) {
	item := models.DocumentType{}
	err := r.db.NewSelect().Model(&item).
		Relation("DocumentTypeFields").
		Where("document_types.id = ?", *id).
		Scan(context.Background())
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.DocumentType{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.DocumentType) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
