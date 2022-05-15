package register

import (
	"context"
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"

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
	Get(string) (*models.Register, error)
	Create(*models.Register) error
	Update(*models.Register) error
	Delete(*string) error

	GetValueTypes() (models.ValueTypes, error)
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Register) error
	getAll(uuid.UUID) ([]*models.Register, error)
	get(string) (*models.Register, error)
	update(*models.Register) error
	delete(*string) error

	getValueTypes() (models.ValueTypes, error)
}

type IFilesService interface {
	Upload(*gin.Context, *models.Register, map[string][]*multipart.FileHeader) error
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
	db          *bun.DB
	ctx         context.Context
	helper      *helper.Helper
	queryFilter *sqlHelper.QueryFilter
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
