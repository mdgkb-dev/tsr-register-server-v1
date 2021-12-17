package users

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.User) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.Users, error) {
	items := make(models.Users, 0)

	err := r.db.NewSelect().
		Model(&items).Scan(r.ctx)
		//Relation("RegistersToUsers")

	return items, err
}

func (r *Repository) get(id *string) (*models.User, error) {
	item := models.User{}
	err := r.db.NewSelect().Model(&item).
		Relation("RegistersUsers").
		Where("users.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.User{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.User) (err error) {
	_, err = r.db.NewUpdate().Model(item).	OmitZero().
		Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) getBySearch(search *string) (models.Users, error) {
	items := make(models.Users, 0)

	err := r.db.NewSelect().
		Model(&items).
		Relation("Human").
		Where("lower(regexp_replace(human.name, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+*search+"%").
		WhereOr("lower(regexp_replace(human.surname, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+*search+"%").
		WhereOr("lower(regexp_replace(human.patronymic, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+*search+"%").
		Scan(r.ctx)
	return items, err
}
