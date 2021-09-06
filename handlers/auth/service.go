package auth

import (
	"errors"
	"fmt"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Register(item *models.User) (*models.TokensWithUser, error) {
	err := item.GenerateHashPassword()

	if err != nil {
		return nil, err
	}
	err = s.repository.create(item)

	token, err := item.CreateToken()
	if err != nil {
		return nil, err
	}
	res := models.TokensWithUser{Token: token, User: item}
	return &res, nil
}

func (s *Service) Login(user *models.User) (*models.TokensWithUser, error) {
	findedUser, err := s.repository.getByLogin(&user.Login)
	if err != nil {
		return nil, err
	}
	fmt.Println("findedUser.ID", findedUser.ID)
	if !findedUser.CompareWithHashPassword(&user.Password) {
		return nil, errors.New("wrong password")
	}
	token, err := findedUser.CreateToken()
	if err != nil {
		return nil, err
	}
	res := models.TokensWithUser{Token: token, User: findedUser}
	return &res, nil
}
