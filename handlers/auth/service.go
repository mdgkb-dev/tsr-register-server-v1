package auth

import (
	"errors"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Register(item *models.User) (*models.TokensWithUser, error) {
	err := item.GenerateHashPassword()

	if err != nil {
		return nil, err
	}
	err = s.repository.create(item)
	if err != nil {
		return nil, err
	}
	ts, err := s.helper.Token.CreateToken(item.ID.String(), "", "")
	if err != nil {
		return nil, err
	}
	return &models.TokensWithUser{Tokens: ts, User: *item}, nil
}

func (s *Service) Login(user *models.User) (*models.TokensWithUser, error) {
	findedUser, err := s.repository.getByLogin(&user.Login)
	if err != nil {
		return nil, err
	}
	if !findedUser.CompareWithHashPassword(user.Password) {
		return nil, errors.New("wrong password")
	}
	ts, err := s.helper.Token.CreateToken(findedUser.ID.String(), "", "")
	if err != nil {
		return nil, err
	}
	return &models.TokensWithUser{Tokens: ts, User: *findedUser}, nil
}

func (s *Service) GetUserByID(id *string) (*models.User, error) {
	return s.repository.getByID(id)
}

func (s *Service) DoesLoginExist(login *string) (bool, error) {
	foundUser, err := s.repository.getByLogin(login)

	if err != nil {
		return false, err
	}

	if foundUser == nil {
		return false, nil
	}

	return true, nil
}
