package contacts

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) Create(c context.Context, item *models.Contact) error {
	_, err := r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (items models.ContactsWithCount, err error) {
	items.Contacts = make(models.Contacts, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&items.Contacts)
	r.helper.SQL.ExtractFTSP(c).HandleQuery(query)
	items.Count, err = query.ScanAndCount(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.Contact, error) {
	item := models.Contact{}
	query := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Where("?TableAlias.id = ?", id)
	err := query.Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Contact{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Contact) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}
