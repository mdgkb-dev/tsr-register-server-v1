package commissionsdrugapplications

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"

	"github.com/uptrace/bun"
)

func (r *Repository) DB() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) SetQueryFilter(c *gin.Context) (err error) {
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Create(item *models.CommissionDrugApplication) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll() (items models.CommissionsDrugApplicationsWithCount, err error) {
	items.CommissionsDrugApplications = make(models.CommissionsDrugApplications, 0)
	query := r.DB().NewSelect().
		Model(&items.CommissionsDrugApplications).
		Relation("CommissionDrugApplicationsDrugApplicationsDoctors.Doctor").
		Relation("PatientDiagnosis.MkbItem").
		Relation("Patient.Human").
		Relation("CommissionDrugApplicationsDrugApplicationstatus")
	items.Count, err = query.ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) Get(id string) (*models.CommissionDrugApplication, error) {
	item := models.CommissionDrugApplication{}
	err := r.DB().NewSelect().
		Model(&item).
		Relation("CommissionDrugApplicationsDrugApplicationsDoctors.Doctor").
		Relation("Patient.Human").
		Where("?TableAlias.id = ?", id).Scan(r.ctx)
	if err != nil {
		return nil, err
	}
	return &item, err
}

func (r *Repository) Delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.CommissionDrugApplication{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(item *models.CommissionDrugApplication) (err error) {
	_, err = r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
