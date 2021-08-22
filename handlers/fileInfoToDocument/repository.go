package fileInfoForDocument

import (
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items []*models.FileInfoToDocument) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.Disability)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.FileInfoToDocument) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("parameter1 = EXCLUDED.parameter1").
		Set("parameter2 = EXCLUDED.parameter2").
		Set("parameter3 = EXCLUDED.parameter3").
		Model(&items).
		Exec(r.ctx)
	return err
}
