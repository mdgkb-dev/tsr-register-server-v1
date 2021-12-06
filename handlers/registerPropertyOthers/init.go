package registerPropertyOthers

import (
	"context"
	"github.com/google/uuid"
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
}

type IService interface {
	GetAll() (models.RegisterPropertyOthers, error)
	Get(*string) (*models.RegisterPropertyOther, error)
	Create(*models.RegisterPropertyOther) error
	Update(*models.RegisterPropertyOther) error
	Delete(*string) error

	UpsertMany(models.RegisterPropertyOthers) error
	DeleteMany([]uuid.UUID) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.RegisterPropertyOther) error
	getAll() ([]*models.RegisterPropertyOther, error)
	get(*string) (*models.RegisterPropertyOther, error)
	update(*models.RegisterPropertyOther) error
	delete(*string) error

	upsertMany(models.RegisterPropertyOthers) error
	deleteMany([]uuid.UUID) error
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

func CreateService(db *bun.DB) *Service {
	repo := NewRepository(db)
	return NewService(repo)
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
