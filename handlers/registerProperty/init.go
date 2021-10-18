package registerProperty

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context) error
	Get(c *gin.Context) error
	Create(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
	GetValueTypes(c *gin.Context) error
}

type IService interface {
	GetAll(*string) ([]*models.RegisterProperty, error)
	Get(*string) (*models.RegisterProperty, error)
	Create(*models.RegisterProperty) error
	Update(*models.RegisterProperty) error
	Delete(*string) error
	GetValueTypes() ([]*models.ValueType, error)
}

type IRepository interface {
	create(*models.RegisterProperty) error
	getAll(*string) ([]*models.RegisterProperty, error)
	get(*string) (*models.RegisterProperty, error)
	update(*models.RegisterProperty) error
	delete(*string) error
	getValueTypes() ([]*models.ValueType, error)
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
