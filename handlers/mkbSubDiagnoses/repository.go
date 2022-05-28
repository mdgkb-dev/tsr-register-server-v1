package mkbSubDiagnoses

import (
	"github.com/gin-gonic/gin"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (items models.MkbSubDiagnoses, err error) {
	err = r.db.NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.MkbSubDiagnosis, error) {
	item := models.MkbSubDiagnosis{}
	err := r.db.NewSelect().
		Model(&item).
		Relation("MkbDiagnosis.MkbGroup").
		Relation("MkbConcreteDiagnosis").
		Where("?TableName.id = ?", id).
		Scan(r.ctx)
	return &item, err
}
