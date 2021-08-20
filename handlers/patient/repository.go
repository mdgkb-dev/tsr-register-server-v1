package patient

import "mdgkb/tsr-tegister-server-v1/models"

func (r *Repository) create(item *models.Patient) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items []*models.Patient, err error) {
	err = r.db.NewSelect().
		Model(&items).
		Relation("AnthropometryData.Anthropometry").
		Relation("Disabilities.Edvs").
		Relation("Human.Documents.DocumentType").
		Relation("RepresentativeToPatient.Representative.Human").
		Relation("RepresentativeToPatient.RepresentativeType").
		Relation("PatientDiagnosis.MkbDiagnosis").
		Relation("PatientDiagnosis.MkbSubDiagnosis").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Patient, error) {
	item := models.Patient{}
	err := r.db.NewSelect().Model(&item).Where("patient.id = ?", *id).Scan(r.ctx)
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
