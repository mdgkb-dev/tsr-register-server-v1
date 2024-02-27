package representativesdomains

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/pro-assistance/pro-assister/middleware"

	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) Create(c context.Context, item *models.RepresentativeDomain) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll(c context.Context) (item models.RepresentativesDomainsWithCount, err error) {
	item.RepresentativesDomains = make(models.RepresentativesDomains, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&item.RepresentativesDomains)

	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) Get(c context.Context, slug string) (*models.RepresentativeDomain, error) {
	item := models.RepresentativeDomain{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Where("?TableAlias.id = ?", slug).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) RepresentativeInDomain(c context.Context, RepresentativeID string) (bool, error) {
	return r.helper.DB.IDB(c).NewSelect().Model((*models.RepresentativeDomain)(nil)).
		Where("?TableAlias.representative_id = ?", RepresentativeID).
		Where("?TableAlias.domain_id in (?)", bun.In(middleware.ClaimDomainIDS.FromContextSlice(c))).
		Exists(r.ctx)
}

func (r *Repository) AddToDomain(c context.Context, items models.RepresentativesDomains) error {
	_, err := r.helper.DB.IDB(c).NewInsert().Model(&items).Exec(c)
	return err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.RepresentativeDomain{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(c context.Context, item *models.RepresentativeDomain) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
