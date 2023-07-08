package commissions

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"

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

func (r *Repository) Create(item *models.Commission) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll() (items models.CommissionsWithCount, err error) {
	items.Commissions = make(models.Commissions, 0)
	query := r.DB().NewSelect().
		Model(&items.Commissions).
		Relation("CommissionsDrugApplications.DrugApplication").
		Relation("CommissionsDoctors.Doctor").
		Relation("PatientDiagnosis.MkbItem").
		Relation("Patient.Human").
		Relation("Patient.PatientDiagnosis.MkbItem").
		Relation("CommissionStatus").
		Relation("DrugRecipe.Drug").
		Relation("DrugRecipe.DrugForm").
		Relation("DrugRecipe.DrugDoze")
	r.queryFilter.HandleQuery(query)
	items.Count, err = query.ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) Get(id string) (*models.Commission, error) {
	item := models.Commission{}
	err := r.DB().NewSelect().
		Model(&item).
		Relation("CommissionsDoctors.Doctor").
		Relation("Patient.Human").
		Relation("DrugRecipe.Drug").
		Relation("DrugRecipe.DrugForm").
		Relation("DrugRecipe.DrugDoze").
		Where("?TableAlias.id = ?", id).Scan(r.ctx)
	if err != nil {
		return nil, err
	}
	return &item, err
}

func (r *Repository) Delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.Commission{}).Where("id = ?", id).Exec(r.ctx)
	return err
}
func (r *Repository) Update(item *models.Commission) (err error) {
	_, err = r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
