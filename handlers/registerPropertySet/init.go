package registerPropertySet

import (
"context"
"github.com/gin-gonic/gin"
"github.com/uptrace/bun"
"mdgkb/tsr-tegister-server-v1/models"
)

type IHandler interface {
GetAll(c *gin.Context) error
Get(c *gin.Context) error
Create(c *gin.Context) error
Update(c *gin.Context) error
Delete(c *gin.Context) error
}

type IService interface {
GetAll() ([]*models.RegisterPropertySet, error)
Get(*string) (*models.RegisterPropertySet, error)
Create(*models.RegisterPropertySet) error
Update(*models.RegisterPropertySet) error
Delete(*string) error
}

type IRepository interface {
create(*models.RegisterPropertySet) error
getAll() ([]*models.RegisterPropertySet, error)
get(*string) (*models.RegisterPropertySet, error)
update(*models.RegisterPropertySet) error
delete(*string) error
}

type Handler struct {
service IService
}

type Service struct {
repository IRepository
}

type Repository struct {
db  *bun.DB
ctx context.Context
}

func CreateHandler(db *bun.DB) *Handler {
repo := NewRepository(db)
service := NewService(repo)
return NewHandler(service)
}

// NewHandler constructor
func NewHandler(s IService) *Handler {
return &Handler{service: s}
}

func NewService(repository IRepository) *Service {
return &Service{repository: repository}
}

func NewRepository(db *bun.DB) *Repository {
return &Repository{db: db, ctx: context.Background()}
}