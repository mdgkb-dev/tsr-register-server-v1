package hmfsescaletestresults

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.HmfseScaleTestResults) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.HmfseScaleTestResult)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.HmfseScaleTestResults) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("hmfse_scale_test_id = EXCLUDED.hmfse_scale_test_id").
		Set("hmfse_scale_question_score_id = EXCLUDED.hmfse_scale_question_score_id").
		Model(&items).
		Exec(r.ctx)
	return err
}
