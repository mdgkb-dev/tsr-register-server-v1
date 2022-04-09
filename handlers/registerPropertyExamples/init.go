package registerPropertyExamples

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
	GetAll() ([]*models.RegisterPropertyExample, error)
	Get(*string) (*models.RegisterPropertyExample, error)
	Create(*models.RegisterPropertyExample) error
	Update(*models.RegisterPropertyExample) error
	Delete(*string) error

	UpsertMany(models.RegisterPropertyExamples) error
	DeleteMany([]uuid.UUID) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.RegisterPropertyExample) error
	getAll() ([]*models.RegisterPropertyExample, error)
	get(*string) (*models.RegisterPropertyExample, error)
	update(*models.RegisterPropertyExample) error
	delete(*string) error

	upsertMany(models.RegisterPropertyExamples) error
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
