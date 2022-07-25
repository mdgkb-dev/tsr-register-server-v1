package mkbgroups

import (
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

func (r *Repository) getAll() (items models.MkbGroups, err error) {
	err = r.db().NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.MkbGroup, error) {
	item := models.MkbGroup{}
	err := r.db().NewSelect().Model(&item).
		Relation("MkbDiagnosis.MkbSubDiagnosis.MkbConcreteDiagnosis").
		Where("?TableName.id = ?", id).Scan(r.ctx)
	return &item, err
}
