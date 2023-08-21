package researches

import (
	"context"

	"mdgkb/tsr-tegister-server-v1/models"
	"mime/multipart"

	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"

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
	GetXlsx(c *gin.Context)
}

type IService interface {
	setQueryFilter(*gin.Context) error
	GetAll() (models.Researches, error)
	Get(string) (*models.Research, error)
	Create(*models.Research) error
	Update(*models.Research) error
	Delete(*string) error

	GetValueTypes() (models.ValueTypes, error)
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	db() *bun.DB
	create(*models.Research) error
	getAll() (models.Researches, error)
	get(string) (*models.Research, error)
	update(*models.Research) error
	delete(*string) error

	getValueTypes() (models.ValueTypes, error)
}

type IFilesService interface {
	Upload(*gin.Context, *models.Research, map[string][]*multipart.FileHeader) error
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
	ctx         context.Context
	helper      *helper.Helper
	queryFilter *sqlHelper.QueryFilter
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

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
