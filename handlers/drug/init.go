package drug

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
	GetAll([]uuid.UUID) ([]*models.Drug, error)
	Get(*string) (*models.Drug, error)
	Create(*models.Drug) error
	Update(*models.Drug) error
	Delete(*string) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Drug) error
	getAll([]uuid.UUID) ([]*models.Drug, error)
	get(*string) (*models.Drug, error)
	update(*models.Drug) error
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
