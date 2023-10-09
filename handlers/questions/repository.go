package questions

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/middleware"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/sqlHelper"
	"github.com/uptrace/bun"
)

func (r *Repository) Create(c context.Context, item *models.Question) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll(c context.Context) (items models.QuestionsWithCount, err error) {
	items.Questions = make(models.Questions, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&items.Questions).
		Relation("AnswerVariants")

	query.Join("join questions_domains qd on qd.question_id = questions.id and qd.domain_id in (?)", bun.In(middleware.ClaimDomainIDS.FromContextSlice(c)))

	i, ok := c.Value("fq").(*sqlHelper.QueryFilter)
	if ok {
		i.HandleQuery(query)
	}
	items.Count, err = query.ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.Question, error) {
	item := models.Question{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("AnswerVariant").
		Where("register_property.id = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) GetAnthropometryQuestions(c context.Context) (models.Questions, error) {
	items := make(models.Questions, 0)
	err := r.helper.DB.IDB(c).NewSelect().Model(&items).
		Where("?TableAlias.code in (?)", bun.In([]string{string(models.AnthropomethryKeyWeight), string(models.AnthropomethryKeyHeight)})).Scan(r.ctx)
	return items, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Question{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Question) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(c context.Context, items models.Questions) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
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

func (r *Repository) deleteMany(c context.Context, idPool []uuid.UUID) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().
		Model((*models.Question)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
