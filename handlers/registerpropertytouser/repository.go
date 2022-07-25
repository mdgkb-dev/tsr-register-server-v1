package registerpropertytouser

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.RegisterPropertyToUser) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(item *models.RegisterPropertyToUser) (err error) {
	_, err = r.db().NewDelete().Model(item).Where("register_property_id = ?", item.RegisterPropertyID).Where("user_id = ?", item.UserID).Exec(r.ctx)
	return err
}
