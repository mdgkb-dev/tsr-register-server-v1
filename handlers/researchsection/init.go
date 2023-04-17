package researchsection

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type IService interface {
	GetAll() ([]*models.Research, error)
	Get(*string) (*models.Research, error)
	Create(*models.Research) error
	Update(*models.Research) error
	Delete(*string) error

	UpsertMany(models.Researches) error
	DeleteMany([]uuid.UUID) error
}

type IRepository interface {
	db() *bun.DB
	create(*models.Research) error
	getAll() ([]*models.Research, error)
	get(*string) (*models.Research, error)
	update(*models.Research) error
	delete(*string) error

	upsertMany(models.Researches) error
	deleteMany([]uuid.UUID) error
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
