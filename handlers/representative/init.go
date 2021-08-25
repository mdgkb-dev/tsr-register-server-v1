package representative

import (
	"context"
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
}

type IService interface {
	GetAll(pagination *httpHelper.Pagination) ([]*models.Representative, error)
	Get(*string) (*models.Representative, error)
	Create(*models.Representative) error
	Update(*models.Representative) error
	Delete(*string) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Representative) error
	getAll(pagination *httpHelper.Pagination) ([]*models.Representative, error)
	get(*string) (*models.Representative, error)
	update(*models.Representative) error
	delete(*string) error

	getOnlyNames() ([]*models.Representative, error)
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
