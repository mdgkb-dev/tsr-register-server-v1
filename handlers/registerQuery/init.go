package registerQuery

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/helpers"
	httpHelper2 "mdgkb/tsr-tegister-server-v1/helpers/httpHelperV2"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Execute(c *gin.Context)
}

type IService interface {
	Create(*models.RegisterQuery) error
	GetAll() (models.RegisterQueries, error)
	Get(*string) (*models.RegisterQuery, error)
	Update(*models.RegisterQuery) error
	Delete(*string) error
	Execute(string) ([]map[string]interface{}, error)
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.RegisterQuery) error
	getAll() (models.RegisterQueries, error)
	get(*string) (*models.RegisterQuery, error)
	update(*models.RegisterQuery) error
	delete(*string) error
	execute(*models.RegisterQuery) ([]map[string]interface{}, error)
}

type Handler struct {
	service IService
	helper  *helpers.Helper
}

type Service struct {
	repository IRepository
	helper     *helpers.Helper
}

type Repository struct {
	db          *bun.DB
	ctx         context.Context
	helper      *helpers.Helper
	queryFilter *httpHelper2.QueryFilter
}

func CreateHandler(db *bun.DB, helper *helpers.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	return NewHandler(service, helper)
}

// NewHandler constructor
func NewHandler(s IService, helper *helpers.Helper) *Handler {
	return &Handler{service: s, helper: helper}
}

func NewService(repository IRepository, helper *helpers.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helpers.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}
