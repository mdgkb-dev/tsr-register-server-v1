package patients

import (
	"mdgkb/tsr-tegister-server-v1/handlers/history"
	"mdgkb/tsr-tegister-server-v1/handlers/human"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *HistoryService) Create(item *models.Patient, requestType models.RequestType) error {
	historyItem := models.History{RequestType: &requestType}
	err := history.CreateService(s.repository.getDB(), s.helper).Create(&historyItem)
	if err != nil {
		return err
	}
	humanHistory := models.HumanHistory{Human: *item.Human}
	err = human.CreateHistoryService(s.repository.getDB()).Create(&humanHistory)
	if err != nil {
		return err
	}
	patientHistory := models.PatientHistory{Patient: *item, HistoryID: historyItem.ID, HumanHistoryID: humanHistory.HumanHistoryID}
	err = s.repository.create(&patientHistory)
	if err != nil {
		return err
	}
	return nil
}

func (s *HistoryService) GetAll(id *string) ([]*models.PatientHistory, error) {
	items, err := s.repository.getAll(id)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *HistoryService) Get(id *string) (*models.PatientHistory, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}
