package insurancecompanytohuman

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items []*models.InsuranceCompanyToHuman) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.InsuranceCompanyToHuman)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)

	return err
}

func (r *Repository) upsertMany(items []*models.InsuranceCompanyToHuman) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("insurance_company_id = EXCLUDED.insurance_company_id").
		Set(`"number" = EXCLUDED."number"`).
		Model(&items).
		Exec(r.ctx)
	return err
}