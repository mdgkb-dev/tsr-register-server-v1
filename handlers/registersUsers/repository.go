package registersUsers

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items models.RegistersUsers) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.RegistersUsers)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.RegistersUsers) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("register_id = EXCLUDED.register_id").
		Set("user_id = EXCLUDED.user_id").
		Model(&items).
		Exec(r.ctx)
	return err
}
