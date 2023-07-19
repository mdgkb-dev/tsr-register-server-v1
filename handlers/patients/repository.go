package patients

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func (r *Repository) DB() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) SetQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}

	r.accessDetails, err = r.helper.Token.GetAccessDetail(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Create(item *models.Patient) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll() (items models.PatientsWithCount, err error) {
	items.Patients = make(models.Patients, 0)
	query := r.DB().NewSelect().
		Model(&items.Patients).
		Relation("Disabilities").
		Relation("Disabilities.Edvs").
		Relation("PatientDrugRegimen").
		Relation("Disabilities.Edvs").
		Relation("Disabilities").
		Relation("Human.Documents.DocumentType").
		Relation("Human.Documents.FileInfoToDocument.FileInfo").
		Relation("Human.Contact").
		Relation("Human.InsuranceCompanyToHuman").
		Relation("PatientsRepresentatives.Representative.Human.Contact").
		Relation("PatientsRepresentatives.RepresentativeType").
		Relation("PatientDiagnosis.MkbItem").
		Relation("PatientsRegisters.Register").
		Relation("PatientsRegisters.User").
		Relation("CreatedBy").
		Relation("UpdatedBy")
	if r.accessDetails != nil && r.accessDetails.UserDomainID != "" {
		query.Where("?TableAlias.domain_id = ?", r.accessDetails.UserDomainID)
	}
	r.queryFilter.HandleQuery(query)
	items.Count, err = query.ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) Get(id string) (*models.Patient, error) {
	item := models.Patient{}
	query := r.DB().NewSelect().Model(&item).
		Relation("Disabilities").
		Relation("Disabilities.Edvs").
		Relation("Disabilities.Edvs.FileInfo").
		Relation("Human.Documents.DocumentType").
		Relation("Human.Documents.FileInfoToDocument.FileInfo").
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
	err := query.Scan(r.ctx)
	return &item, err
}

func (r *Repository) Delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.Patient{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(item *models.Patient) (err error) {
	_, err = r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
