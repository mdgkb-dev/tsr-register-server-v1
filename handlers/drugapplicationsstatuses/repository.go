package drugapplicationsstatuses

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
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

func (r *Repository) Create(item *models.DrugApplicationStatus) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll() (item models.DrugApplicationsStatusesWithCount, err error) {
	item.DrugApplicationsStatuses = make(models.DrugApplicationsStatuses, 0)
	query := r.DB().NewSelect().Model(&item.DrugApplicationsStatuses)

	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) Get(slug string) (*models.DrugApplicationStatus, error) {
	item := models.DrugApplicationStatus{}
	err := r.DB().NewSelect().Model(&item).
		Relation("DrugApplicationStatus").
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) Delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.DrugApplicationStatus{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(item *models.DrugApplicationStatus) (err error) {
	_, err = r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
