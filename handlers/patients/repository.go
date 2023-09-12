package patients

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/middleware"
	"mdgkb/tsr-tegister-server-v1/models"

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
		Relation("UpdatedBy")

	query.Join("join patients_domains on patients_domains.patient_id = patients_view.id and patients_domains.domain_id in (?)", bun.In(middleware.ClaimDomainIDS.FromContextSlice(c)))
	r.helper.SQL.ExtractQueryFilter(c).HandleQuery(query)
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
		Relation("Anamneses").
		Where("?TableAlias.id = ?", id)
	err := query.Scan(c)
	return &item, err
}

func (r *Repository) GetBySnilsNumber(c context.Context, snillsNumber string) (*models.Patient, error) {
	item := models.Patient{}
	query := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Human").
		Join("join humans h on h.id = ?TableAlias.human_id").
		Join("join documents d on d.human_id = h.id").
		Join("join document_field_values dfv on dfv.document_id = d.id and dfv.value_string = ?", snillsNumber).
		Join("join document_types dt on d.document_type_id = dt.id and dt.code = ?", models.DocumentTypeCodeSnils)
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

func (r *Repository) GetForExport(c context.Context, idPool []string) (items models.PatientsWithCount, err error) {
	items.Patients = make(models.Patients, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&items.Patients).
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
		Relation("Anamneses")
	if len(idPool) > 0 {
		query = query.Where("?TableAlias.id in (?)", bun.In(idPool))
	}

	query.Join("join patients_domains on patients_domains.patient_id = patients_view.id and patients_domains.domain_id in (?)", bun.In(middleware.ClaimDomainIDS.FromContextSlice(c)))
	err = query.Scan(c)
	return items, err
}
