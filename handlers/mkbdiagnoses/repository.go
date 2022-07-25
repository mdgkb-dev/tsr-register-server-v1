package mkbdiagnoses

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

func (r *Repository) getAll() (items models.MkbDiagnoses, err error) {
	err = r.db().NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.MkbDiagnosis, error) {
	item := models.MkbDiagnosis{}
	err := r.db().NewSelect().
		Model(&item).
		Relation("MkbGroup").
		Relation("MkbSubDiagnosis.MkbConcreteDiagnosis").
		Where("?TableName.id = ?", id).Scan(r.ctx)
	return &item, err
}
