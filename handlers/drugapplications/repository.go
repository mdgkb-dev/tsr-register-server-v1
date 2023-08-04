package drugapplications

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

func (r *Repository) Create(item *models.DrugApplication) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll() (items models.DrugApplicationsWithCount, err error) {
	items.DrugApplications = make(models.DrugApplications, 0)
	query := r.DB().NewSelect().
		Model(&items.DrugApplications).
		Relation("DrugApplicationStatus").
		Relation("FundContract.DrugArrives", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("drug_arrives.stage")
		}).
		Relation("FundContract.DrugArrives.DrugDecreases").
		Relation("DrugApplicationFiles.FileInfo").
		Relation("CommissionsDrugApplications.Commission.Patient.Human")
	r.queryFilter.HandleQuery(query)
	items.Count, err = query.ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) Get(id string) (*models.DrugApplication, error) {
	item := models.DrugApplication{}
	err := r.DB().NewSelect().
		Model(&item).
		Relation("DrugApplicationStatus").
		Where("?TableAlias.id = ?", id).Scan(r.ctx)
	if err != nil {
		return nil, err
	}
	return &item, err
}

func (r *Repository) Delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.DrugApplication{}).Where("id = ?", id).Exec(r.ctx)
	return err
}
func (r *Repository) Update(item *models.DrugApplication) (err error) {
	_, err = r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
