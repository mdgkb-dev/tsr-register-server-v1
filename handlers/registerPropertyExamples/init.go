package registerPropertyExamples

import (
	"context"
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"
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
	GetAll() ([]*models.RegisterPropertyExample, error)
	Get(*string) (*models.RegisterPropertyExample, error)
	Create(*models.RegisterPropertyExample) error
	Update(*models.RegisterPropertyExample) error
	Delete(*string) error

	UpsertMany(models.RegisterPropertyExamples) error
	DeleteMany([]uuid.UUID) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.RegisterPropertyExample) error
	getAll() ([]*models.RegisterPropertyExample, error)
	get(*string) (*models.RegisterPropertyExample, error)
	update(*models.RegisterPropertyExample) error
	delete(*string) error

	upsertMany(models.RegisterPropertyExamples) error
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
