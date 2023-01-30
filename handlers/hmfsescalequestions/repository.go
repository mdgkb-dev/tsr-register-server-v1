package hmfsescalequestions

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.HmfseScaleQuestion) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.HmfseScaleQuestions, error) {
	items := make(models.HmfseScaleQuestions, 0)
	err := r.db().NewSelect().Model(&items).
		Relation("HmfseScaleQuestionScores").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.HmfseScaleQuestion, error) {
	item := models.HmfseScaleQuestion{}
	err := r.db().NewSelect().Model(&item).Where("id = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.HmfseScaleQuestion{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.HmfseScaleQuestion) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
