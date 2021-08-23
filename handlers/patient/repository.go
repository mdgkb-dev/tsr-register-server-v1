package patient

import (
	"fmt"
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Patient) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll(offset *int) (items []*models.Patient, err error) {
	err = r.db.NewSelect().
		Model(&items).
		Relation("AnthropometryData.Anthropometry").
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
		Offset(*offset).
		Limit(25).
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Patient, error) {
	item := models.Patient{}
	err := r.db.NewSelect().Model(&item).
		Relation("AnthropometryData.Anthropometry").
		Relation("Disabilities.Edvs.Period").
		Relation("Disabilities.Edvs.FileInfo").
		Relation("Disabilities.Period").
		Relation("Human.Documents.DocumentType").
		Relation("Human.Documents.FileInfoToDocument.FileInfo").
		Relation("Human.Documents.DocumentFieldValues.DocumentTypeField").
		Relation("Human.InsuranceCompanyToHuman.InsuranceCompany").
		Relation("Human.Contact").
		Relation("RepresentativeToPatient.Representative.Human").
		Relation("RepresentativeToPatient.RepresentativeType").
		Relation("PatientDiagnosis.MkbDiagnosis.MkbGroup").
		Relation("PatientDiagnosis.MkbDiagnosis.MkbSubDiagnosis").
		Relation("PatientDiagnosis.MkbSubDiagnosis").
		Relation("RegisterToPatient.Register").
		Where("patient.id = ?", *id).Scan(r.ctx)
	fmt.Println(item.PatientDiagnosis[0].MkbSubDiagnosisID)
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
