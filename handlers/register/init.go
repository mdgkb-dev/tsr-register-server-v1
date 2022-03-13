package register

import (
	"context"
	"github.com/google/uuid"
	"mdgkb/tsr-tegister-server-v1/helpers"
	"mdgkb/tsr-tegister-server-v1/helpers/httpHelper"
	httpHelper2 "mdgkb/tsr-tegister-server-v1/helpers/httpHelperV2"
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

	GetValueTypes(c *gin.Context)
	GetFlatXlsx(c *gin.Context)
}

type IService interface {
	GetAll(uuid.UUID) ([]*models.Register, error)
	Get(*httpHelper.QueryFilter) (*models.Register, error)
	Create(*models.Register) error
	Update(*models.Register) error
	Delete(*string) error

	GetValueTypes() (models.ValueTypes, error)
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Register) error
	getAll(uuid.UUID) ([]*models.Register, error)
	get(*httpHelper.QueryFilter) (*models.Register, error)
	update(*models.Register) error
	delete(*string) error

	getValueTypes() (models.ValueTypes, error)
}

type IFilesService interface {
	Upload(*gin.Context, *models.Register, map[string][]*multipart.FileHeader) error
}

type Handler struct {
	service      IService
	helper       *helpers.Helper
	filesService IFilesService
}

type Service struct {
	repository IRepository
	helper     *helpers.Helper
}

type Repository struct {
	db          *bun.DB
	ctx         context.Context
	helper      *helpers.Helper
	queryFilter *httpHelper2.QueryFilter
}

type FilesService struct {
	helper *helpers.Helper
}

func CreateHandler(db *bun.DB, helper *helpers.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service, filesService, helper)
}

// NewHandler constructor
func NewHandler(s IService, filesService IFilesService, helper *helpers.Helper) *Handler {
	return &Handler{service: s, filesService: filesService, helper: helper}
}

func NewService(repository IRepository, helper *helpers.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helpers.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}

func NewFilesService(helper *helpers.Helper) *FilesService {
	return &FilesService{helper: helper}
}
