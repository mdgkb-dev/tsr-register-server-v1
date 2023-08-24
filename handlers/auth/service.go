package auth

import (
	"errors"
	"mdgkb/tsr-tegister-server-v1/handlers/users"
	"mdgkb/tsr-tegister-server-v1/handlers/usersaccounts"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Register(item *models.UserAccount) (t *models.TokensWithUser, err error) {
	err = usersaccounts.CreateService(s.helper).Create(item)
	if err != nil {
		return nil, err
	}
	user := models.User{}
	user.UserAccountID = item.ID
	user.UserAccount = item

	err = users.CreateService(s.helper).Create(&user)
	if err != nil {
		return nil, err
	}

	token, err := s.helper.Token.CreateToken(&user)
	if err != nil {
		return nil, err
	}
	t.Init(token, user)
	return t, err
}

func (s *Service) Login(item *models.UserAccount, skipPassword bool) (t *models.TokensWithUser, err error) {
	foundedAccount, err := usersaccounts.CreateService(s.helper).GetByEmail(item.Email)
	if err != nil {
		return nil, err
	}
	if !foundedAccount.CompareWithHashPassword(item.Password) && !skipPassword {
		return nil, errors.New("wrong email or password")
	}
	user, err := users.CreateService(s.helper).GetByUserAccountID(foundedAccount.ID.UUID.String())
	if err != nil {
		return nil, err
	}
	token, err := s.helper.Token.CreateToken(user)
	if err != nil {
		return nil, err
	}
	t = &models.TokensWithUser{}
	t.Init(token, *user)
	return t, err
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
