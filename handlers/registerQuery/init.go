package registerQuery

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
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
	Get(string) (*models.RegisterQuery, error)
	Update(*models.RegisterQuery) error
	Delete(string) error
	Execute(*models.RegisterQuery) error
}

type IRepository interface {
	db() *bun.DB
	create(*models.RegisterQuery) error
	getAll() (models.RegisterQueries, error)
	get(string) (*models.RegisterQuery, error)
	update(*models.RegisterQuery) error
	delete(string) error
	execute(*models.RegisterQuery) error
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
