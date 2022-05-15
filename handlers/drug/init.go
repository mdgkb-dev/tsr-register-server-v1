package drug

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
