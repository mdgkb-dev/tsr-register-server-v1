package mkb

import (
	"fmt"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) getAllClasses() (items models.MkbClasses, err error) {
	err = r.db.NewSelect().Model(&items).Order("mkb_class.number").Scan(r.ctx)
	return items, err
}

func (r *Repository) getGroupsByClassId(classID string) (items models.MkbGroups, err error) {
	items = models.MkbGroups{}
	err = r.db.NewSelect().Model(&items).
		Where("mkb_group.mkb_class_id = ?", classID).
		Order("mkb_group.range_start").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) getSubGroupByGroupId(groupId string) (items models.MkbSubGroups, err error) {
	items = models.MkbSubGroups{}
	err = r.db.NewSelect().Model(&items).
		Where("mkb_sub_group.mkb_group_id = ?", groupId).
		Order("mkb_sub_group.range_start").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) getDiagnosisByClassId(classId string) (items models.MkbDiagnoses, err error) {
	err = r.db.NewSelect().Model(&items).
		Where("mkb_diagnosis.mkb_class_id = ?", classId).
		Where("mkb_diagnosis.mkb_group_id IS NULL").
		Where("mkb_diagnosis.mkb_sub_group_id IS NULL").
		Where("mkb_diagnosis.mkb_sub_sub_group_id IS NULL").
		Order("mkb_diagnosis.code").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) getDiagnosisByGroupId(groupId string) (models.MkbDiagnoses, error) {
	items := make(models.MkbDiagnoses, 0)
	err := r.db.NewSelect().Model(&items).
		Relation("MkbGroup").
		Relation("MkbSubDiagnosis.MkbDiagnosis").
		Relation("MkbSubDiagnosis.MkbConcreteDiagnosis").
		Where("mkb_diagnosis.mkb_group_id = ?", groupId).
		Where("mkb_diagnosis.mkb_sub_group_id IS NULL").
		Where("mkb_diagnosis.mkb_sub_sub_group_id IS NULL").
		Order("mkb_diagnosis.code").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) getDiagnosisBySubGroupId(subGroupId string) (items models.MkbDiagnoses, err error) {
	err = r.db.NewSelect().Model(&items).
		Relation("MkbGroup").
		Where("mkb_diagnosis.mkb_sub_group_id = ?", subGroupId).
		Where("mkb_diagnosis.mkb_sub_sub_group_id IS NULL").
		Order("mkb_diagnosis.code").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) getDiagnosisBySubSubGroupId(subSubGroupId string) (items models.MkbDiagnoses, err error) {
	err = r.db.NewSelect().Model(&items).
		Relation("MkbGroup").
		Where("mkb_diagnosis.mkb_sub_sub_group_id = ?", subSubGroupId).
		Order("mkb_diagnosis.code").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) getSubDiagnosisByDiagnosisId(diagnosisId string) (items models.MkbSubDiagnoses, err error) {
	items = models.MkbSubDiagnoses{}
	err = r.db.NewSelect().Model(&items).
		Where("mkb_sub_diagnosis.mkb_diagnosis_id = ?", diagnosisId).
		Order("mkb_sub_diagnosis.sub_code").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) getGroupsByRange(search string) (models.MkbGroups, error) {
	items := make(models.MkbGroups, 0)
	lenOfSearch := len([]rune(search))
	err := r.db.NewSelect().Model(&items).
		Where("lower(left(mkb_group.range_start, ?)) <= lower(?)", lenOfSearch, search).
		Where("lower(left(mkb_group.range_end, ?)) >= lower(?)", lenOfSearch, search).
		Scan(r.ctx)
	return items, err
}

func (r *Repository) getGroupBySearch(search string) (models.MkbGroups, error) {
	items := make(models.MkbGroups, 0)
	err := r.db.NewSelect().Model(&items).
		Where("lower(regexp_replace(mkb_group.name, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+search+"%").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) getDiagnosisBySearch(search string) (models.MkbDiagnoses, error) {
	items := make(models.MkbDiagnoses, 0)
	err := r.db.NewSelect().Model(&items).
		Relation("MkbGroup").
		Relation("MkbSubDiagnosis.MkbConcreteDiagnosis.MkbSubDiagnosis.MkbDiagnosis.MkbGroup").
		Where("lower(regexp_replace(mkb_diagnosis.name, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+search+"%").
		WhereOr("lower(regexp_replace(mkb_diagnosis.code, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+search+"%").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) getSubDiagnosesBySearch(search string) (models.MkbSubDiagnoses, error) {
	items := make(models.MkbSubDiagnoses, 0)
	err := r.db.NewSelect().Model(&items).
		Relation("MkbDiagnosis.MkbGroup").
		Relation("MkbConcreteDiagnosis").
		Where("lower(regexp_replace(mkb_sub_diagnosis.name, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+search+"%").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) updateRelevant(id, table string) (err error) {
	query := fmt.Sprintf(`UPDATE %s SET relevant = NOT relevant WHERE id = '%s' ;`, table, id)
	_, err = r.db.Exec(query)
	return err
}

func (r *Repository) updateName(id, name, table string) (err error) {
	query := fmt.Sprintf(`UPDATE %s SET "name" = '%s' WHERE id = '%s' ;`, table, name, id)
	_, err = r.db.Exec(query)
	return err
}

func (r *Repository) getConcreteDiagnosisBySubDiagnosisId(diagnosisId string) (items models.MkbConcreteDiagnoses, err error) {
	items = models.MkbConcreteDiagnoses{}
	err = r.db.NewSelect().Model(&items).
		Where("mkb_concrete_diagnosis.mkb_sub_diagnosis_id = ?", diagnosisId).
		Order("mkb_concrete_diagnosis.name").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) getConcreteDiagnosisBySearch(search string) (models.MkbConcreteDiagnoses, error) {
	items := make(models.MkbConcreteDiagnoses, 0)
	err := r.db.NewSelect().Model(&items).
		Relation("MkbSubDiagnosis.MkbDiagnosis.MkbGroup").
		Where("lower(regexp_replace(mkb_concrete_diagnosis.name, '[^а-яА-Яa-zA-Z0-9 ]', '', 'g')) LIKE lower(?)", "%"+search+"%").
		Scan(r.ctx)
	return items, err
}
