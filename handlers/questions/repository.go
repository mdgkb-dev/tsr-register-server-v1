package questions

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.Question) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll(registerID *string) ([]*models.Question, error) {
	items := []*models.Question{}
	query := r.db().NewSelect().Model(&items)

	if *registerID != "" {
		query.
			Join("join register_group on register_group.id = ?TableName.register_group_id ").
			Where("register_group.register_id = ? ", *registerID)
	}

	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Question, error) {
	item := models.Question{}
	err := r.db().NewSelect().Model(&item).
		Relation("AnswerVariant").
		Relation("AnswerVariant").
		Where("register_property.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Question{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Question) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) getValueTypes() ([]*models.ValueType, error) {
	items := []*models.ValueType{}
	err := r.db().NewSelect().
		Model(&items).
		Scan(r.ctx)
	return items, err
}

//
//var existsRegisterPropertyWithGroupId string = `
//exists
//(
//	select *
//	from
//		register_property_to_register_group as rptrg
//	where
//		rptrg.register_property_id = register_property.id
//		and exists
//		(
//			select *
//			from
//				register_group_to_register as rgtr
//			where
//				rgtr.register_group_id = rptrg.register_group_id
//				and rgtr.register_id = ?
//		)
//)`

func (r *Repository) upsertMany(items models.Questions) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set(`name = EXCLUDED.name`).
		Set(`short_name = EXCLUDED.short_name`).
		Set(`with_other = EXCLUDED.with_other`).
		Set(`is_files_storage = EXCLUDED.is_files_storage`).
		Set(`col_width = EXCLUDED.col_width`).
		Set(`register_property_order = EXCLUDED.register_property_order`).
		Set(`value_type_id = EXCLUDED.value_type_id`).
		Set(`register_group_id = EXCLUDED.register_group_id`).
		Set(`tag = EXCLUDED.tag`).
		Set(`age_compare = EXCLUDED.age_compare`).
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.Question)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}