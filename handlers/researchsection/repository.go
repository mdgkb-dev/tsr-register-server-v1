package researchsection

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.Research) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items []*models.Research, err error) {
	err = r.db().NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Research, error) {
	item := models.Research{}
	err := r.db().NewSelect().
		Model(&item).
		Relation("RegisterPropertyToRegisterGroup.Question").
		Where("register_group.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Research{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Research) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Researches) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set(`name = EXCLUDED.name`).
		Set(`register_id = EXCLUDED.register_id`).
		Set(`with_dates = EXCLUDED.with_dates`).
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.Research)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
