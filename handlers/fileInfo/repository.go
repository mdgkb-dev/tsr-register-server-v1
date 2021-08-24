package fileInfo

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) create(item *models.FileInfo) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) get(id *string) (*models.FileInfo, error) {
	item := models.FileInfo{}
	err := r.db.NewSelect().Model(&item).
		Where("file_infos.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) update(item *models.FileInfo) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.FileInfo) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("original_name = EXCLUDED.original_name").
		Set("file_system_path = EXCLUDED.file_system_path").
		Model(item).
		Exec(r.ctx)
	return err
}

func (r *Repository) createMany(items []*models.FileInfo) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items []*models.FileInfo) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("original_name = EXCLUDED.original_name").
		Set("file_system_path = EXCLUDED.file_system_path").
		Model(&items).
		Exec(r.ctx)
	return err
}
