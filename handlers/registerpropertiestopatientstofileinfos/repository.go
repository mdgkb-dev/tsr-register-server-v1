package registerpropertiestopatientstofileinfos

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.RegisterPropertiesToPatientsToFileInfos) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.RegisterPropertyToPatientToFileInfo)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.RegisterPropertiesToPatientsToFileInfos) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("register_property_to_patient_id = EXCLUDED.register_property_to_patient_id").
		Set("file_info_id = EXCLUDED.file_info_id").
		Exec(r.ctx)
	return err
}
