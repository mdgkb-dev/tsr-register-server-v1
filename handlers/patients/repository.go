package patients

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}


func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.HTTP.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) create(item *models.Patient) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items models.PatientsWithCount, err error) {
	query := r.db.NewSelect().
		Model(&items.Patients).
		Relation("HeightWeight").
		Relation("ChestCircumference").
		Relation("HeadCircumference").
		Relation("PatientDrugRegimen").
		Relation("Disabilities.Edvs.Period").
		Relation("Disabilities.Period").
		Relation("Human.Documents.DocumentType").
		Relation("Human.Documents.FileInfoToDocument.FileInfo").
		Relation("Human.Contact").
		Relation("Human.InsuranceCompanyToHuman").
		Relation("RepresentativeToPatient.Representative.Human").
		Relation("RepresentativeToPatient.RepresentativeType").
		Relation("PatientDiagnosis.MkbDiagnosis").
		Relation("PatientDiagnosis.MkbSubDiagnosis").
		Relation("RegisterToPatient.Register").
		Relation("CreatedBy").
		Relation("UpdatedBy").
		Join("JOIN regions_users ON patients.region_id = regions_users.region_id AND regions_users.user_id = ?", r.queryFilter.UserID)

	r.helper.HTTP.CreateWithDeletedQuery(query, r.queryFilter.WithDeleted)
	r.helper.HTTP.CreatePaginationQuery(query, r.queryFilter.Pagination)
	r.helper.HTTP.CreateFilter(query, r.queryFilter.FilterModels)
	r.helper.HTTP.CreateOrder(query, r.queryFilter.SortModels, []string{"human.surname", "human.name"})
	items.Count, err = query.ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) get(id *string, withDeleted bool) (*models.Patient, error) {
	item := models.Patient{}
	query := r.db.NewSelect().Model(&item).
		Relation("HeightWeight").
		Relation("ChestCircumference").
		Relation("HeadCircumference").
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
		Relation("RegisterPropertyOthersPatient").
		Relation("CreatedBy").
		Relation("UpdatedBy").
		Where("patients.id = ?", *id)
	if withDeleted {
		query = query.WhereAllWithDeleted()
	}
	err := query.Scan(r.ctx)
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

func (r *Repository) getOnlyNames() (items models.PatientsWithCount, err error) {
	items.Count, err = r.db.NewSelect().
		Model(&items.Patients).
		Relation("Human").
		Order("human.surname").
		Order("human.name").
		Order("human.patronymic").
		ScanAndCount(r.ctx)
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

func (r *Repository) getDisabilities() (item models.PatientsWithCount, err error) {
	item.Patients = make([]*models.Patient, 0)
	item.Count, err = r.db.NewSelect().
		Model(&item.Patients).
		Join("JOIN disability ON disability.patient_id = patients.id").
		Relation("Human").
		Relation("Disabilities.Period").
		Relation("Disabilities.Edvs.Period").
		Relation("Disabilities.Edvs.FileInfo").
		Group("patients.id", "human.id", "human.name", "human.surname", "human.patronymic", "human.is_male", "human.date_birth", "human.address_registration", "human.address_residential", "human.contact_id", "human.photo_id", "human.deleted_at").
		ScanAndCount(r.ctx)
	return item, err
}
