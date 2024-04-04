package answervariants

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.AnswerVariant) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(context.TODO())
	return err
}

func (r *Repository) getAll() (items []*models.AnswerVariant, err error) {
	err = r.db().NewSelect().Model(&items).Scan(context.TODO())
	return items, err
}

func (r *Repository) get(id *string) (*models.AnswerVariant, error) {
	item := models.AnswerVariant{}
	err := r.db().NewSelect().Model(&item).Where("AnswerVariant.id = ?", *id).Scan(context.TODO())
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.AnswerVariant{}).Where("id = ?", *id).Exec(context.TODO())
	return err
}

func (r *Repository) update(item *models.AnswerVariant) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(context.TODO())
	return err
}

func (r *Repository) upsertMany(items models.AnswerVariants) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set(`name = EXCLUDED.name`).
		Set(`register_property_id = EXCLUDED.register_property_id`).
		Set(`register_property_radio_order = EXCLUDED.register_property_radio_order`).
		Exec(context.TODO())
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.AnswerVariant)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(context.TODO())
	return err
}
