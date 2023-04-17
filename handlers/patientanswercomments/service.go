package patientanswercomments

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) UpsertMany(items models.PatientAnswerComments) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.upsertMany(items)
}
