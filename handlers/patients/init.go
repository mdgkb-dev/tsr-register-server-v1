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
	GetAllHistory(c *gin.Context)
	GetHistory(c *gin.Context)
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
	getDB() *bun.DB
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

type IHistoryRepository interface {
	getDB() *bun.DB
	create(*models.PatientHistory) error
	getAll(*string) ([]*models.PatientHistory, error)
	get(*string) (*models.PatientHistory, error)
}

type IHistoryService interface {
	Create(*models.Patient, models.RequestType) error
	GetAll(*string) ([]*models.PatientHistory, error)
	Get(*string) (*models.PatientHistory, error)
}

type HistoryService struct {
	repository IHistoryRepository
	helper     *helper.Helper
}

type HistoryRepository struct {
	db  *bun.DB
	ctx context.Context
}

type Handler struct {
	service        IService
	filesService   IFilesService
	helper         *helper.Helper
	historyService IHistoryService
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
	repoHistory := NewHistoryRepository(db)
	historyService := NewHistoryService(repoHistory)
	return NewHandler(service, filesService, historyService, helper)
}

// NewHandler constructor
func NewHandler(s IService, filesService IFilesService, historyService IHistoryService, helper *helper.Helper) *Handler {
	return &Handler{service: s, filesService: filesService, helper: helper, historyService: historyService}
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

func NewHistoryService(repository IHistoryRepository) *HistoryService {
	return &HistoryService{repository: repository}
}

func NewHistoryRepository(db *bun.DB) *HistoryRepository {
	return &HistoryRepository{db: db, ctx: context.Background()}
}
