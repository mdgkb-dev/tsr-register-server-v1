package registertopatient

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/helper"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type IHandler interface {
	Delete(c *gin.Context)
}

type IService interface {
	CreateMany([]*models.ResearchResult) error
	Delete(string) error
}

type IRepository interface {
	db() *bun.DB
	createMany([]*models.ResearchResult) error
	upsertMany([]*models.ResearchResult) error
	deleteMany([]uuid.UUID) error
	delete(string) error
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
