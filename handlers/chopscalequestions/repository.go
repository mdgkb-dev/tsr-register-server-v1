package chopscalequestions

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.ChopScaleQuestion) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.ChopScaleQuestions, error) {
	items := make(models.ChopScaleQuestions, 0)
	err := r.db().NewSelect().Model(&items).
		Relation("ChopScaleQuestionScores", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("chop_scale_question_scores.score DESC")
		}).
		Order("item_order").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.ChopScaleQuestion, error) {
	item := models.ChopScaleQuestion{}
	err := r.db().NewSelect().Model(&item).Where("id = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.ChopScaleQuestion{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.ChopScaleQuestion) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
