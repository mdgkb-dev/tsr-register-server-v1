package registerQuery

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	Create(c *gin.Context) error
	GetAll(c *gin.Context) error
	Get(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
}

type IService interface {
	Create(*models.RegisterQuery) error
	GetAll() ([]*models.RegisterQuery, error)
	Get(*string) (*models.RegisterQuery, error)
	Update(*models.RegisterQuery) error
	Delete(*string) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.RegisterQuery) error
	getAll() ([]*models.RegisterQuery, error)
	get(*string) (*models.RegisterQuery, error)
	update(*models.RegisterQuery) error
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

func NewHandler(s IService) *Handler {
	return &Handler{service: s}
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, ctx: context.Background()}
}
