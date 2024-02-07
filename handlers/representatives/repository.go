package representatives

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/middleware"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) Create(c context.Context, item *models.Representative) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll(c context.Context) (items models.RepresentativesWithCount, err error) {
	items.Representatives = make(models.Representatives, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&items.Representatives).
		Relation("Human.Documents.DocumentType").
		Relation("Human.Documents.DocumentFileInfos.FileInfo").
		Relation("Human.Contact").
		Relation("PatientsRepresentatives.Patient.Human").
		Relation("PatientsRepresentatives.RepresentativeType")
	// Order("human.surname")
	query.Join("join representatives_domains on representatives_domains.representative_id = representatives_view.id and representatives_domains.domain_id in (?)", bun.In(middleware.ClaimDomainIDS.FromContextSlice(c)))
	r.helper.SQL.ExtractFTSP(c).HandleQuery(query)
	r.helper.SQL.ExtractQueryFilter(c).HandleQuery(query)
	items.Count, err = query.ScanAndCount(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.Representative, error) {
	item := models.Representative{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Human.Contact").
		Relation("Human.Photo").
		Relation("PatientsRepresentatives.Patient.Human").
		Relation("PatientsRepresentatives.RepresentativeType").
		Where("?TableAlias.id = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Representative{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Representative) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) GetBySnilsNumber(c context.Context, snillsNumber string) (*models.Representative, error) {
	item := models.Representative{}
	query := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Human").
		Join("join humans h on h.id = ?TableAlias.human_id").
		Join("join documents d on d.human_id = h.id").
		Join("join document_field_values dfv on dfv.document_id = d.id and dfv.value_string = ?", snillsNumber).
		Join("join document_types dt on d.document_type_id = dt.id and dt.code = ?", models.DocumentTypeCodeSnils)
	err := query.Scan(c)
	return &item, err
}
