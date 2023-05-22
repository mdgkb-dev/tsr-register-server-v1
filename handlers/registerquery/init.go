package registerquery

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/pro-assistance/pro-assister/helper"

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
	Create(*models.ResearchQuery) error
	GetAll() (models.ResearchQueries, error)
	Get(string) (*models.ResearchQuery, error)
	Update(*models.ResearchQuery) error
	Delete(string) error
	Execute(*models.ResearchQuery) error
}

type IRepository interface {
	db() *bun.DB
	create(*models.ResearchQuery) error
	getAll() (models.ResearchQueries, error)
	get(string) (*models.ResearchQuery, error)
	update(*models.ResearchQuery) error
	delete(string) error
	execute(*models.ResearchQuery) error
}

type Handler struct {
	service IService
	helper  *helper.Helper
}

type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
	service := NewService(repo, helper)
	return NewHandler(service, helper)
}

// NewHandler constructor
func NewHandler(s IService, helper *helper.Helper) *Handler {
	return &Handler{service: s, helper: helper}
}

func CreateService(helper *helper.Helper) *Service {
	repo := NewRepository(helper)
	return NewService(repo, helper)
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}
