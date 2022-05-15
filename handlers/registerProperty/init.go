package registerProperty

import (
	"context"
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetValueTypes(c *gin.Context)
}

type IService interface {
	GetAll(*string) ([]*models.RegisterProperty, error)
	Get(*string) (*models.RegisterProperty, error)
	Create(*models.RegisterProperty) error
	Update(*models.RegisterProperty) error
	Delete(*string) error
	GetValueTypes() ([]*models.ValueType, error)

	UpsertMany(models.RegisterProperties) error
	DeleteMany([]uuid.UUID) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.RegisterProperty) error
	getAll(*string) ([]*models.RegisterProperty, error)
	get(*string) (*models.RegisterProperty, error)
	update(*models.RegisterProperty) error
	delete(*string) error
	getValueTypes() ([]*models.ValueType, error)

	upsertMany(models.RegisterProperties) error
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
	db     *bun.DB
	ctx    context.Context
	helper *helper.Helper
}

func CreateHandler(db *bun.DB, helper *helper.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	return NewHandler(service, helper)
}

func CreateService(db *bun.DB, helper *helper.Helper) *Service {
	repo := NewRepository(db, helper)
	return NewService(repo, helper)
}

// NewHandler constructor
func NewHandler(s IService, helper *helper.Helper) *Handler {
	return &Handler{service: s, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helper.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}
