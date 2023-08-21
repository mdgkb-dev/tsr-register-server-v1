package patients

import (
	"context"
	"fmt"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/pro-assistance/pro-assister/sqlHelper"
	"github.com/uptrace/bun"
)

func (r *Repository) Create(c context.Context, item *models.Patient) error {
	_, err := r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (items models.PatientsWithCount, err error) {
	items.Patients = make(models.Patients, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&items.Patients).
		Relation("Disabilities").
		Relation("Disabilities.Edvs").
		Relation("PatientDrugRegimen").
		Relation("Disabilities.Edvs").
		Relation("Disabilities").
		Relation("Human.Documents.DocumentType").
		Relation("Human.Documents.DocumentFileInfos.FileInfo").
		Relation("Human.Contact").
		Relation("Human.InsuranceCompanyToHuman").
		Relation("PatientsRepresentatives.Representative.Human.Contact").
		Relation("PatientsRepresentatives.RepresentativeType").
		Relation("PatientDiagnosis.MkbItem").
		Relation("PatientsRegisters.Register").
		Relation("PatientsRegisters.User").
		Relation("CreatedBy").
		Relation("UpdatedBy").Limit(20)
	//userDomainId, err := r.helper.Token.ExtractTokenMetadata(c.Request, "user_id")
	//if userDomainId != "" {
	//	query.Where("?TableAlias.domain_id = ?")
	//}
	i, ok := c.Value(fqKey{}).(*sqlHelper.QueryFilter)
	if !ok {
		fmt.Println("1", i)
	}
	//fmt.Println(r.helper.SQL.ExtractQueryFilter(c))
	i.HandleQuery(query)
	items.Count, err = query.ScanAndCount(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.Patient, error) {
	item := models.Patient{}
	query := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Disabilities").
		Relation("Disabilities.Edvs").
		Relation("Disabilities.Edvs.FileInfo").
		Relation("Human.Documents.DocumentType").
		Relation("Human.Documents.DocumentFileInfos.FileInfo").
		Relation("Human.Documents.DocumentFieldValues.DocumentTypeField").
		Relation("Human.InsuranceCompanyToHuman.InsuranceCompany").
		Relation("Human.Contact").
		Relation("Human.Photo").
		Relation("PatientsRepresentatives.Representative.Human.Contact").
		Relation("PatientsRepresentatives.RepresentativeType").
		Relation("PatientDiagnosis.MkbItem").
		Relation("PatientDiagnosis.Anamneses").
		Relation("PatientsResearchesPools.ResearchesPool").
		Relation("PatientsResearches.ResearchResults.Answers.SelectedAnswerVariants").
		Relation("PatientsRegisters.Register").
		Relation("PatientsRegisters.User").
		Relation("Commissions", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("commissions.number")
		}).
		Relation("Commissions.CommissionsDoctors.Doctor").
		Relation("Commissions.PatientDiagnosis.MkbItem").
		Where("?TableAlias.id = ?", id)
	err := query.Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Patient{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Patient) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}
