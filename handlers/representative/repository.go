package representative

import (
	"github.com/gin-gonic/gin"
	"mdgkb/tsr-tegister-server-v1/models"

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

func (r *Repository) create(item *models.Representative) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items models.RepresentativesWithCount, err error) {
	query := r.db().NewSelect().
		Model(&items.Representatives).
		Relation("Human.Documents.DocumentType").
		Relation("Human.Documents.FileInfoToDocument.FileInfo").
		Relation("Human.Contact").
		Relation("RepresentativeToPatient.Patient.Human").
		Relation("RepresentativeToPatient.RepresentativeType").
		Order("human.surname")
	r.queryFilter.HandleQuery(query)
	//httpHelper.CreateFilter(query, queryFilter.FilterModels)
	//httpHelper.CreatePaginationQuery(query, queryFilter.Pagination)
	items.Count, err = query.ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) getOnlyNames() (items models.RepresentativesWithCount, err error) {
	items.Count, err = r.db().NewSelect().
		Model(&items.Representatives).
		Relation("Human").
		Order("human.surname").
		Order("human.name").
		Order("human.patronymic").
		ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Representative, error) {
	item := models.Representative{}
	err := r.db().NewSelect().Model(&item).
		Relation("Human.Documents.DocumentType").
		Relation("Human.Documents.FileInfoToDocument.FileInfo").
		Relation("Human.Documents.DocumentFieldValues.DocumentTypeField").
		Relation("Human.Contact").
		Relation("Human.Photo").
		Relation("RepresentativeToPatient.Patient.Human").
		Relation("RepresentativeToPatient.RepresentativeType").
		Where("?TableAlias.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Representative{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Representative) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) getBySearch(search *string) (models.Representatives, error) {
	items := make(models.Representatives, 0)

	err := r.db().NewSelect().
		Model(&items).
		Relation("Human").
		Where("lower(regexp_replace(human.name, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+*search+"%").
		WhereOr("lower(regexp_replace(human.surname, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+*search+"%").
		WhereOr("lower(regexp_replace(human.patronymic, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+*search+"%").
		Scan(r.ctx)
	return items, err
}
