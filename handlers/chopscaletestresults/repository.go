package chopscaletestresults

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.ChopScaleTestResults) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.ChopScaleTestResult)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.ChopScaleTestResults) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("chop_scale_test_id = EXCLUDED.chop_scale_test_id").
		Set("chop_scale_question_score_id = EXCLUDED.chop_scale_question_score_id").
		Model(&items).
		Exec(r.ctx)
	return err
}
