package usersaccounts

import (
	"errors"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.UserAccount) error {
	err := item.HashPassword()
	if err != nil {
		return err
	}
	return s.repository.Create(item)
}

func (s *Service) CheckAccountPassword(item *models.UserAccount, skipPassword bool) error {
	foundUser, err := s.repository.GetByEmail(item.Email)
	if err != nil {
		return errors.New("wrong email or password")
	}
	if !foundUser.CompareWithHashPassword(item.Password) && !skipPassword {
		return errors.New("wrong email or password")
	}
	return nil
}

func (s *Service) Get(id string) (*models.UserAccount, error) {
	return s.repository.Get(id)
}

func (s *Service) GetByEmail(email string) (*models.UserAccount, error) {
	return s.repository.GetByEmail(email)
}

func (s *Service) UpdateUUID(id string) error {
	return s.repository.UpdateUUID(id)
}

func (s *Service) UpdatePassword(item *models.UserAccount) error {
	err := item.HashPassword()
	if err != nil {
		return err
	}
	return s.repository.UpdatePassword(item)
}


