package researches

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/middleware"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) create(item *models.Research) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll(c context.Context) (items models.Researches, err error) {
	query := r.db().NewSelect().
		Model(&items).
		Relation("Questions", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("questions.item_order")
		}).
		Relation("Questions.AnswerVariants", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("answer_variants.item_order")
		}).
		Relation("Questions.QuestionExamples").
		Relation("Questions.ValueType").
		Relation("Questions.QuestionVariants", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("question_variants.name")
		}).
		Relation("Questions.Children.ValueType").
		Relation("Questions.Children.AnswerVariants").
		Relation("Formulas.FormulaResults")

	query.Join("join researches_domains on researches_domains.research_id = researches.id and researches_domains.domain_id in (?)", bun.In(middleware.ClaimDomainIDS.FromContextSlice(c)))
	r.helper.SQL.ExtractQueryFilter(c).HandleQuery(query)
	err = query.Scan(r.ctx)

	return items, err
}

func (r *Repository) get(id string) (*models.Research, error) {
	item := models.Research{}
	err := r.db().NewSelect().
		Model(&item).
		Relation("Questions", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("questions.item_order")
		}).
		Relation("Questions.AnswerVariants", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("answer_variants.item_order")
		}).
		Relation("Questions.QuestionExamples").
		Relation("Questions.ValueType").
		Relation("Questions.QuestionVariants", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("question_variants.name")
		}).
		Relation("Questions.Children", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("questions.item_order")
		}).
		Relation("Questions.Children.Children", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("questions.item_order")
		}).
		Relation("Questions.Children.ValueType").
		Relation("Questions.Children.Children.ValueType").
		Relation("Questions.Children.AnswerVariants", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("answer_variants.item_order")
		}).
		Relation("Formulas.FormulaResults").
		Where("?TableAlias.id = ?", id).Scan(r.ctx)
	if err != nil {
		return nil, err
	}
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

func (r *Repository) getValueTypes() (models.ValueTypes, error) {
	items := make(models.ValueTypes, 0)
	err := r.db().NewSelect().
		Model(&items).
		Scan(r.ctx)
	return items, err
}

func (r *Repository) GetForExport(c context.Context, idPool []string) (items models.Researches, err error) {
	query := r.db().NewSelect().
		Model(&items).
		Relation("Questions", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("questions.item_order")
		}).
		Relation("Questions.AnswerVariants", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("answer_variants.item_order")
		}).
		Relation("Questions.QuestionExamples").
		Relation("Questions.ValueType").
		Relation("Questions.QuestionVariants", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("question_variants.name")
		}).
		Relation("Questions.Children.ValueType").
		Relation("Questions.Children.AnswerVariants").
		Relation("Formulas.FormulaResults")

	if len(idPool) > 0 {
		query = query.Where("?TableAlias.id in (?)", bun.In(idPool))
	}

	query.Join("join researches_domains on researches_domains.research_id = researches.id and researches_domains.domain_id in (?)", bun.In(middleware.ClaimDomainIDS.FromContextSlice(c)))
	r.helper.SQL.ExtractQueryFilter(c).HandleQuery(query)
	err = query.Scan(r.ctx)
	return items, err
}
