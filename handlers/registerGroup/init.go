package registerGroup

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
	GetAll() ([]*models.RegisterGroup, error)
	Get(*string) (*models.RegisterGroup, error)
	Create(*models.RegisterGroup) error
	Update(*models.RegisterGroup) error
	Delete(*string) error

	UpsertMany(models.RegisterGroups) error
	DeleteMany([]uuid.UUID) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.RegisterGroup) error
	getAll() ([]*models.RegisterGroup, error)
	get(*string) (*models.RegisterGroup, error)
	update(*models.RegisterGroup) error
	delete(*string) error

	upsertMany(models.RegisterGroups) error
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
