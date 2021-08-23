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
		Set("file_info_id = EXCLUDED.file_info_id").
		Set("document_id = EXCLUDED.document_id").
		Model(&items).
		Exec(r.ctx)
	return err
}
