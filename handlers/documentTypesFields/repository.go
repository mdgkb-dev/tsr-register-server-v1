package documentTypeFields

import (
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) createMany(items []*models.DocumentTypeField) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.DocumentTypeField)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.DocumentTypeField) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("name = EXCLUDED.name").
		Set(`"order" = EXCLUDED."order"`).
		Set("type = EXCLUDED.type").
		Exec(r.ctx)
	return err
}

//
//func (r *Repository) update(item *models.DocumentTypeField) (err error) {
//	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
//	return err
//}
