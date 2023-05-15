package patients

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
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

func (r *Repository) create(item *models.Patient) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items models.PatientsWithCount, err error) {
	query := r.db().NewSelect().
		Model(&items.Patients).
		Relation("HeightWeight").
		Relation("ChestCircumference").
		Relation("HeadCircumference").
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
	//Join("JOIN regions_users ON patients.region_id = regions_users.region_id AND regions_users.user_id = ?")
	r.queryFilter.HandleQuery(query)
	items.Count, err = query.ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) get(id *string, withDeleted bool) (*models.Patient, error) {
	item := models.Patient{}
	query := r.db().NewSelect().Model(&item).
		//Relation("HeightWeight").
		//Relation("ChestCircumference").
		//Relation("HeadCircumference").
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
		//Relation("PatientHistories.User").

		//Relation("PatientDiagnosis.Anamnesis").
		//Relation("ResearchResult.ResearchesPool").
		//Relation("RegisterGroupsToPatient.Answer.Question").
		//Relation("RegisterGroupsToPatient.Answer.RegisterPropertiesToPatientsToFileInfos.FileInfo").
		//Relation("RegisterGroupsToPatient.Answer.AnswerVariant").
		//Relation("RegisterGroupsToPatient.PatientAnswerComments").
		//Relation("PatientDrugRegimen.DrugRegimen.Drug").
		//Relation("PatientDrugRegimen.PatientDrugRegimenItems").
		//Relation("ChopScaleTests.ChopScaleTestResults.ChopScaleQuestionScore").
		//Relation("HmfseScaleTests.HmfseScaleTestResults.HmfseScaleQuestionScore").
		//Relation("CreatedBy").
		//Relation("UpdatedBy").
		Where("?TableAlias.id = ?", *id)
	if withDeleted {
		query = query.WhereAllWithDeleted()
	}
	err := query.Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Patient{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Patient) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) getOnlyNames() (items models.PatientsWithCount, err error) {
	items.Count, err = r.db().NewSelect().
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

	err := r.db().NewSelect().
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
	item.Count, err = r.db().NewSelect().
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
