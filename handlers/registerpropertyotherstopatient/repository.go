package registerpropertyotherstopatient

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) upsertMany(items models.RegisterPropertyOthersToPatient) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set(`value = EXCLUDED.value`).
		Exec(r.ctx)
	return err
}
