package mkb

import (
	"fmt"
	"mdgkb/tsr-tegister-server-v1/models"
)

type CompositionMkb struct {
	MkbGroups    models.MkbGroups    `json:"mkbGroups"`
	MkbSubGroups models.MkbSubGroups `json:"mkbSubGroups"`
	MkbDiagnosis models.MkbDiagnoses `json:"mkbDiagnosis"`
}

func (s *Service) GetAllClasses() (models.MkbClasses, error) {
	classes, err := s.repository.getAllClasses()
	if err != nil {
		return nil, err
	}
	return classes, nil
}

func (s *Service) GetGroupByClassId(classID string) (*CompositionMkb, error) {
	groups, err := s.repository.getGroupsByClassId(classID)
	if err != nil {
		return nil, err
	}
	diagnosis, err := s.repository.getDiagnosisByClassId(classID)
	if err != nil {
		return nil, err
	}
	res := CompositionMkb{MkbGroups: groups, MkbDiagnosis: diagnosis}
	return &res, nil
}

func (s *Service) GetGroupChildrens(groupId string) (*CompositionMkb, error) {
	subGroups, err := s.repository.getSubGroupByGroupId(groupId)
	if err != nil {
		return nil, err
	}
	diagnosis, err := s.repository.getDiagnosisByGroupId(groupId)
	if err != nil {
		return nil, err
	}
	fmt.Println(subGroups)

	res := CompositionMkb{MkbSubGroups: subGroups, MkbDiagnosis: diagnosis}
	fmt.Println(res.MkbSubGroups)
	return &res, nil
}

func (s *Service) GetDiagnosisByGroupId(groupId string) (models.MkbDiagnoses, error) {
	diagnosis, err := s.repository.getDiagnosisByGroupId(groupId)
	if err != nil {
		return nil, err
	}
	return diagnosis, nil
}

func (s *Service) GetSubGroupChildrens(groupId string) (*CompositionMkb, error) {
	diagnosis, err := s.repository.getDiagnosisBySubGroupId(groupId)
	if err != nil {
		return nil, err
	}
	res := CompositionMkb{MkbDiagnosis: diagnosis}
	return &res, nil
}

func (s *Service) GetSubDiagnosisByDiagnosisId(diagnosisId string) (models.MkbSubDiagnoses, error) {
	return s.repository.getSubDiagnosisByDiagnosisId(diagnosisId)
}

func (s *Service) GetDiagnosisBySearch(search string) (models.MkbDiagnoses, error) {
	diagnosis, err := s.repository.getDiagnosisBySearch(search)
	if err != nil {
		return nil, err
	}
	return diagnosis, nil
}

func (s *Service) GetSubDiagnosesBySearch(search string) (models.MkbSubDiagnoses, error) {
	diagnosis, err := s.repository.getSubDiagnosesBySearch(search)
	if err != nil {
		return nil, err
	}
	return diagnosis, nil
}

func (s *Service) GetGroupsBySearch(search string) (models.MkbGroups, error) {
	lenOfSearch := len([]rune(search))
	if lenOfSearch < 3 {
		groups, err := s.repository.getGroupsByRange(search)
		if err != nil {
			return nil, err
		}
		return groups, nil
	}
	groups, err := s.repository.getGroupBySearch(search)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (s *Service) UpdateRelevant(id, model string) error {
	table := getTableName(model)
	err := s.repository.updateRelevant(id, table)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateName(id, name, model string) error {
	table := getTableName(model)
	err := s.repository.updateName(id, name, table)
	if err != nil {
		return err
	}
	return nil
}

func getTableName(model string) string {
	tableName := ""
	switch model {
	case "MkbCLass":
		tableName = "mkb_class"
	case "MkbGroup":
		tableName = "mkb_group"
	case "MkbSubGroup":
		tableName = "mkb_sub_groups"
	case "MkbSubSubGroup":
		tableName = "mkb_sub_sub_groups"
	case "MkbDiagnosis":
		tableName = "mkb_diagnosis"
	case "MkbSubDiagnosis":
		tableName = "mkb_sub_diagnosis"
	}
	return tableName
}

func (s *Service) GetConcreteDiagnosisBySearch(search string) (models.MkbConcreteDiagnoses, error) {
	return s.repository.getConcreteDiagnosisBySearch(search)
}
func (s *Service) GetConcreteDiagnosisBySubDiagnosisId(diagnosisId string) (models.MkbConcreteDiagnoses, error) {
	return s.repository.getConcreteDiagnosisBySubDiagnosisId(diagnosisId)
}
