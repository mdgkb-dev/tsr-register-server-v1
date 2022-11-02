package fileinfos

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IHandler interface {
	Download(c *gin.Context)
}

type IService interface {
	Create(*models.FileInfo) error
	Get(*string) (*models.FileInfo, error)
	Update(*models.FileInfo) error
	Upsert(*models.FileInfo) error
	Delete(uuid.NullUUID) error

	CreateMany([]*models.FileInfo) error
	UpsertMany([]*models.FileInfo) error
}

type IRepository interface {
	create(*models.FileInfo) error
	get(*string) (*models.FileInfo, error)
	update(*models.FileInfo) error
	upsert(*models.FileInfo) error
	delete(uuid.NullUUID) error

	createMany([]*models.FileInfo) error
	upsertMany([]*models.FileInfo) error
}
type IFilesService interface {
	GetFullPath(*string) *string
}

type Handler struct {
	service      IService
	helper       *helper.Helper
	filesService IFilesService
}

type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

type FilesService struct {
	helper *helper.Helper
}

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service, filesService, helper)
}

// NewHandler constructor
func NewHandler(s IService, filesService IFilesService, helper *helper.Helper) *Handler {
	return &Handler{service: s, filesService: filesService, helper: helper}
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

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
