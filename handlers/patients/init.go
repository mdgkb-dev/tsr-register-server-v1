package patients

import (
	"context"

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
}

type IService interface {
	setQueryFilter(*gin.Context) error

	GetAll() (models.PatientsWithCount, error)
	GetOnlyNames() (models.PatientsWithCount, error)
	Get(*string, bool) (*models.Patient, error)
	Create(*models.Patient) error
	Update(*models.Patient) error
	Delete(*string) error

	GetBySearch(*string) ([]*models.Patient, error)
	GetDisabilities() (models.PatientsWithCount, error)
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	db() *bun.DB
	create(*models.Patient) error
	getAll() (models.PatientsWithCount, error)
	get(*string, bool) (*models.Patient, error)
	update(*models.Patient) error
	delete(*string) error

	getOnlyNames() (models.PatientsWithCount, error)
	getBySearch(*string) ([]*models.Patient, error)
	getDisabilities() (models.PatientsWithCount, error)
}

type IFilesService interface {
	Upload(*gin.Context, *models.Patient, map[string][]*multipart.FileHeader) error
}

type Handler struct {
	service      IService
	filesService IFilesService
	helper       *helper.Helper
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
