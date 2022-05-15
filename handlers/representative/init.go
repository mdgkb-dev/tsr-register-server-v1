package representative

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/tsr-tegister-server-v1/models"
	"mime/multipart"

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
	GetAll() (models.RepresentativesWithCount, error)
	GetOnlyNames() (models.RepresentativesWithCount, error)
	Get(*string) (*models.Representative, error)
	Create(*models.Representative) error
	Update(*models.Representative) error
	Delete(*string) error

	GetBySearch(*string) ([]*models.Representative, error)
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Representative) error
	getAll() (models.RepresentativesWithCount, error)
	get(*string) (*models.Representative, error)
	update(*models.Representative) error
	delete(*string) error

	getOnlyNames() (models.RepresentativesWithCount, error)
	getBySearch(*string) (models.Representatives, error)
}

type IFilesService interface {
	Upload(*gin.Context, *models.Representative, map[string][]*multipart.FileHeader) error
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
	db     *bun.DB
	ctx    context.Context
	helper *helper.Helper
}

type FilesService struct {
	helper *helper.Helper
}

func CreateHandler(db *bun.DB, helper *helper.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service, filesService, helper)
}

// NewHandler constructor
func NewHandler(s IService, filesService IFilesService, helper *helper.Helper) *Handler {
	return &Handler{service: s, filesService: filesService, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helper.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
