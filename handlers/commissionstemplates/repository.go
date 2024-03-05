package commissionstemplates

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) create(item *models.CommissionTemplate) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items models.CommissionsTemplates, err error) {
	err = r.db().NewSelect().
		Model(&items).
		Relation("CommissionsDoctorsTemplates.Doctor").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.CommissionTemplate, error) {
	item := models.CommissionTemplate{}
	err := r.db().NewSelect().
		Model(&item).
		Relation("CommissionsTemplatesDoctors.Doctor").
		Where("?TableAlias.id = ?", id).Scan(r.ctx)
	if err != nil {
		return nil, err
	}
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.CommissionTemplate{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.CommissionTemplate) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
