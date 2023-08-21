package questions

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) DB() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) SetQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Create(item *models.Question) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll() (items models.QuestionsWithCount, err error) {
	items.Questions = make(models.Questions, 0)
	query := r.DB().NewSelect().
		Model(&items.Questions).
		Relation("AnswerVariants")

	r.queryFilter.HandleQuery(query)
	items.Count, err = query.ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) Get(id string) (*models.Question, error) {
	item := models.Question{}
	err := r.DB().NewSelect().Model(&item).
		Relation("AnswerVariant").
		Where("register_property.id = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) Delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.Question{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(item *models.Question) (err error) {
	_, err = r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Questions) (err error) {
	_, err = r.DB().NewInsert().On("conflict (id) do update").
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
	_, err = r.DB().NewDelete().
		Model((*models.Question)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
