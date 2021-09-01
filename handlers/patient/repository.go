package patient

import (
	"mdgkb/tsr-tegister-server-v1/helpers/httpHelper"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Patient) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll(pagination *httpHelper.Pagination) (items []*models.Patient, err error) {
	err = r.db.NewSelect().
		Model(&items).
		Relation("HeightWeight").
		Relation("PatientDrugRegimen").
		Relation("Disabilities.Edvs.Period").
		Relation("Disabilities.Period").
		Relation("Human.Documents.DocumentType").
		Relation("Human.Documents.FileInfoToDocument.FileInfo").
		Relation("Human.Contact").
		Relation("RepresentativeToPatient.Representative.Human").
		Relation("RepresentativeToPatient.RepresentativeType").
		Relation("PatientDiagnosis.MkbDiagnosis").
		Relation("PatientDiagnosis.MkbSubDiagnosis").
		Relation("RegisterToPatient.Register").
		Order("human.surname").
		Order("human.name").
		Offset(*pagination.Offset).
		Limit(*pagination.Limit).
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Patient, error) {
	item := models.Patient{}
	err := r.db.NewSelect().Model(&item).
		Relation("HeightWeight").
		Relation("Disabilities.Period").
		Relation("Disabilities.Edvs.Period").
		Relation("Disabilities.Edvs.FileInfo").
		Relation("Human.Documents.DocumentType").
		Relation("Human.Documents.FileInfoToDocument.FileInfo").
		Relation("Human.Documents.DocumentFieldValues.DocumentTypeField").
		Relation("Human.InsuranceCompanyToHuman.InsuranceCompany").
		Relation("Human.Contact").
		Relation("Human.Photo").
		Relation("RepresentativeToPatient.Representative.Human").
		Relation("RepresentativeToPatient.RepresentativeType").
		Relation("PatientDiagnosis.MkbDiagnosis.MkbGroup").
		Relation("PatientDiagnosis.MkbDiagnosis.MkbSubDiagnosis").
		Relation("PatientDiagnosis.MkbSubDiagnosis").
		Relation("PatientDiagnosis.PatientDiagnosisAnamnesis").
		Relation("RegisterToPatient.Register").
		Relation("RegisterPropertyToPatient.RegisterProperty").
		Relation("RegisterPropertySetToPatient.RegisterPropertySet").
		Relation("PatientDrugRegimen.DrugRegimen.Drug").
		Relation("PatientDrugRegimen.PatientDrugRegimenItems").
		Where("patient.id = ?", *id).Scan(r.ctx)

	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Patient{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Patient) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) getOnlyNames() (items []*models.Patient, err error) {
	err = r.db.NewSelect().
		Model(&items).
		Relation("Human").
		Order("human.surname").
		Order("human.name").
		Order("human.patronymic").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) getBySearch(search *string) ([]*models.Patient, error) {
	items := make([]*models.Patient, 0)

	err := r.db.NewSelect().
		Model(&items).
		Relation("Human").
		Where("lower(regexp_replace(human.name, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+*search+"%").
		WhereOr("lower(regexp_replace(human.surname, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+*search+"%").
		WhereOr("lower(regexp_replace(human.patronymic, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+*search+"%").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) getDisabilities() ([]*models.Patient, error) {
	items := make([]*models.Patient, 0)
	err := r.db.NewSelect().
		Model(&items).
		Join("JOIN disability ON disability.patient_id = patient.id").
		Relation("Human").
		Relation("Disabilities.Period").
		Relation("Disabilities.Edvs.Period").
		Relation("Disabilities.Edvs.FileInfo").
		Group("patient.id", "human.id").
		Scan(r.ctx)
	return items, err
}
