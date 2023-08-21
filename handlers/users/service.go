package users

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
)

func (s *Service) Create(item *models.User) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(item *models.User) error {
	//err := human.CreateService(s.helper).Upsert(item.Human)
	//if err != nil {
	//	return err
	//}
	//item.UUID.UUID, err = uuid.NewUUID()
	//if err != nil {
	//	return err
	//}
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	return nil
}

func (s *Service) Upsert(item *models.User) error {
	//item.UUID.UUID, err = uuid.NewUUID()
	//if err != nil {
	//	return err
	//}
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	return nil
}

func (s *Service) UpsertEmail(item *models.User) error {
	err := s.repository.upsertEmail(item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() (models.UsersWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.User, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) GetByUserAccountID(userAccountID string) (*models.User, error) {
	item, err := s.repository.getByUserAccountID(userAccountID)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) AddToUser(values map[string]interface{}, table string) error {
	err := s.repository.addToUser(values, table)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) RemoveFromUser(values map[string]interface{}, table string) error {
	err := s.repository.removeFromUser(values, table)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) setQueryFilter(c *gin.Context) (err error) {
	err = s.repository.setQueryFilter(c)
	return err
}
