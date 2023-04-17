package auth

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *ValidateService) Login(item *models.Login) error {
	return s.helper.Validator.Validate(item)
}
