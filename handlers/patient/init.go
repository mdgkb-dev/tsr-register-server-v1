package patient

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/helpers/httpHelper"
	"mdgkb/tsr-tegister-server-v1/helpers/uploadHelper"
	"mdgkb/tsr-tegister-server-v1/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context) error
	Get(c *gin.Context) error
	Create(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
	GetAllHistory(c *gin.Context) error
	GetHistory(c *gin.Context) error
	GetAllWithDeleted(c *gin.Context) error
}

type IService interface {
	GetAll(filter *httpHelper.QueryFilter) (models.PatientsWithCount, error)
	GetOnlyNames() (models.PatientsWithCount, error)
	Get(*string, bool) (*models.Patient, error)
	Create(*models.Patient) error
	Update(*models.Patient) error
	Delete(*string) error

	GetBySearch(*string) ([]*models.Patient, error)
	GetDisabilities() (models.PatientsWithCount, error)
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.Patient) error
	getAll(*httpHelper.QueryFilter) (models.PatientsWithCount, error)
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
}

type HistoryRepository struct {
	db  *bun.DB
	ctx context.Context
}

type Handler struct {
	service        IService
	historyService IHistoryService
	filesService   IFilesService
}

type Service struct {
	repository IRepository
}

type Repository struct {
	db  *bun.DB
	ctx context.Context
}

type FilesService struct {
	uploader uploadHelper.Uploader
}

func CreateHandler(db *bun.DB, uploader *uploadHelper.Uploader) *Handler {
	repo := NewRepository(db)
	service := NewService(repo)
	repoHistory := NewHistoryRepository(db)
	historyService := NewHistoryService(repoHistory)
	filesService := NewFilesService(uploader)
	return NewHandler(service, filesService, historyService)
}

func NewHandler(service IService, filesService IFilesService, historyService IHistoryService) *Handler {
	return &Handler{service: service, filesService: filesService, historyService: historyService}
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, ctx: context.Background()}
}

func NewFilesService(uploader *uploadHelper.Uploader) *FilesService {
	return &FilesService{uploader: *uploader}
}

func NewHistoryService(repository IHistoryRepository) *HistoryService {
	return &HistoryService{repository: repository}
}

func NewHistoryRepository(db *bun.DB) *HistoryRepository {
	return &HistoryRepository{db: db, ctx: context.Background()}
}
