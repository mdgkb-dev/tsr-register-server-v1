package auth

import (
	"mdgkb/tsr-tegister-server-v1/models"

	_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) getByLogin(login *string) (*models.User, error) {
	user := models.User{}
	err := r.db.NewSelect().Model(&user).
		Relation("RegisterPropertyToUser").
		Where("login = ?", *login).
		Scan(r.ctx)
	return &user, err
}

func (r *Repository) getByID(id *string) (*models.User, error) {
	user := models.User{}
	err := r.db.NewSelect().Model(&user).
		Relation("RegisterPropertyToUser").
		Where("id = ?", *id).
		Scan(r.ctx)
	return &user, err
}

func (r *Repository) create(user *models.User) (err error) {
	_, err = r.db.NewInsert().Model(user).Exec(r.ctx)
	return err
}
