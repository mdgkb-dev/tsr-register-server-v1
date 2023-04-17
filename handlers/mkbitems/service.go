package mkbitems

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

//type CompositionMkb struct {
//	MkbGroups    models.MkbGroups    `json:"mkbGroups"`
//	MkbSubGroups models.MkbSubGroups `json:"mkbSubGroups"`
//	MkbDiagnosis models.MkbDiagnoses `json:"mkbDiagnosis"`
//}

//func (s *Service) GetAllClasses() (models.MkbClasses, error) {
//	classes, err := s.repository.getAllClasses()
//	if err != nil {
//		return nil, err
//	}
//	return classes, nil
//}
//
//func (s *Service) GetGroupByClassID(classID string) (*CompositionMkb, error) {
//	groups, err := s.repository.getGroupsByClassID(classID)
//	if err != nil {
//		return nil, err
//	}
//	diagnosis, err := s.repository.getDiagnosisByClassID(classID)
//	if err != nil {
//		return nil, err
//	}
//	res := CompositionMkb{MkbGroups: groups, MkbDiagnosis: diagnosis}
//	return &res, nil
//}
//
//func (s *Service) GetGroupChildrens(groupID string) (*CompositionMkb, error) {
//	subGroups, err := s.repository.getSubGroupByGroupID(groupID)
//	if err != nil {
//		return nil, err
//	}
//	diagnosis, err := s.repository.getDiagnosisByGroupID(groupID)
//	if err != nil {
//		return nil, err
//	}
//	fmt.Println(subGroups)
//
//	res := CompositionMkb{MkbSubGroups: subGroups, MkbDiagnosis: diagnosis}
//	fmt.Println(res.MkbSubGroups)
//	return &res, nil
//}
//
//func (s *Service) GetDiagnosisByGroupID(groupID string) (models.MkbDiagnoses, error) {
//	diagnosis, err := s.repository.getDiagnosisByGroupID(groupID)
//	if err != nil {
//		return nil, err
//	}
//	return diagnosis, nil
//}
//
//func (s *Service) GetSubGroupChildrens(groupID string) (*CompositionMkb, error) {
//	diagnosis, err := s.repository.getDiagnosisBySubGroupID(groupID)
//	if err != nil {
//		return nil, err
//	}
//	res := CompositionMkb{MkbDiagnosis: diagnosis}
//	return &res, nil
//}
//
//func (s *Service) GetSubDiagnosisByDiagnosisID(diagnosisID string) (models.MkbSubDiagnoses, error) {
//	return s.repository.getSubDiagnosisByDiagnosisID(diagnosisID)
//}
//
//func (s *Service) GetDiagnosisBySearch(search string) (models.MkbDiagnoses, error) {
//	diagnosis, err := s.repository.getDiagnosisBySearch(search)
//	if err != nil {
//		return nil, err
//	}
//	return diagnosis, nil
//}
//
//func (s *Service) GetSubDiagnosesBySearch(search string) (models.MkbSubDiagnoses, error) {
//	diagnosis, err := s.repository.getSubDiagnosesBySearch(search)
//	if err != nil {
//		return nil, err
//	}
//	return diagnosis, nil
//}
//
//func (s *Service) GetGroupsBySearch(search string) (models.MkbGroups, error) {
//	lenOfSearch := len([]rune(search))
//	if lenOfSearch < 3 {
//		groups, err := s.repository.getGroupsByRange(search)
//		if err != nil {
//			return nil, err
//		}
//		return groups, nil
//	}
//	groups, err := s.repository.getGroupBySearch(search)
//	if err != nil {
//		return nil, err
//	}
//	return groups, nil
//}
//
//func (s *Service) UpdateRelevant(id, model string) error {
//	table := getTableName(model)
//	err := s.repository.updateRelevant(id, table)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (s *Service) UpdateName(id, name, model string) error {
//	table := getTableName(model)
//	err := s.repository.updateName(id, name, table)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func getTableName(model string) string {
//	tableName := ""
//	switch model {
//	case "MkbClass":
//		tableName = "mkb_class"
//	case "MkbGroup":
//		tableName = "mkb_groups"
//	case "MkbSubGroup":
//		tableName = "mkb_sub_groups"
//	case "MkbSubSubGroup":
//		tableName = "mkb_sub_sub_groups"
//	case "MkbDiagnosis":
//		tableName = "mkb_diagnosis"
//	case "MkbSubDiagnosis":
//		tableName = "mkb_sub_diagnosis"
//	}
//	return tableName
//}
//
//func (s *Service) GetConcreteDiagnosisBySearch(search string) (models.MkbConcreteDiagnoses, error) {
//	return s.repository.getConcreteDiagnosisBySearch(search)
//}
//func (s *Service) GetConcreteDiagnosisBySubDiagnosisID(diagnosisID string) (models.MkbConcreteDiagnoses, error) {
//	return s.repository.getConcreteDiagnosisBySubDiagnosisID(diagnosisID)
//}
//
//func (s *Service) SelectMkbElement(id string) (*models.MkbClass, *models.MkbElement, error) {
//	element, err := s.repository.selectMkbElement(id)
//	if err != nil {
//		return nil, nil, err
//	}
//	class, err := s.repository.getMkbClass(element.ClassID)
//	if err != nil {
//		return nil, nil, err
//	}
//	return class, element, nil
//}

func (s *Service) GetTree() (*models.MkbItem, error) {
	items, err := s.repository.GetTree()
	if err != nil {
		return nil, err
	}
	root := items.GetRoot()
	models.BuildTree(root, items)
	return root, nil
}

func (s *Service) Get(id string) (*models.MkbItem, error) {
	item, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}
