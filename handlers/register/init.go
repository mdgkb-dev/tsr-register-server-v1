package register

import (
	"context"
	"github.com/google/uuid"
	"mdgkb/tsr-tegister-server-v1/helpers/httpHelper"
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
	GetAll(uuid.UUID) ([]*models.Register, error)
	Get(*httpHelper.QueryFilter) (*models.Register, error)
	Create(*models.Register) error
	Update(*models.Register) error
	Delete(*string) error

	GetValueTypes() (models.ValueTypes, error)
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Register) error
	getAll(uuid.UUID) ([]*models.Register, error)
	get(*httpHelper.QueryFilter) (*models.Register, error)
	update(*models.Register) error
	delete(*string) error

	getValueTypes() (models.ValueTypes, error)
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
