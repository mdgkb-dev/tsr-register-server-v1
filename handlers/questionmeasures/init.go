package questionmeasures

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"

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
	GetAll() ([]*models.QuestionMeasure, error)
	Get(*string) (*models.QuestionMeasure, error)
	Create(*models.QuestionMeasure) error
	Update(*models.QuestionMeasure) error
	Delete(*string) error

	UpsertMany(models.QuestionMeasures) error
	DeleteMany([]uuid.UUID) error
}

type IRepository interface {
	db() *bun.DB
	create(*models.QuestionMeasure) error
	getAll() ([]*models.QuestionMeasure, error)
	get(*string) (*models.QuestionMeasure, error)
	update(*models.QuestionMeasure) error
	delete(*string) error

	upsertMany(models.QuestionMeasures) error
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

func CreateService(helper *helper.Helper) *Service {
	repo := NewRepository(helper)
	return NewService(repo, helper)
}

// NewHandler constructor
func NewHandler(s IService, helper *helper.Helper) *Handler {
	return &Handler{service: s, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}
