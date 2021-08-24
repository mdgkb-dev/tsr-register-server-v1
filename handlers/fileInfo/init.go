package fileInfo

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/helpers/uploadHelper"
	"mdgkb/tsr-tegister-server-v1/models"
)

type IHandler interface {
	Download(c *gin.Context) error
}

type IService interface {
	Create(*models.FileInfo) error
	Get(*string) (*models.FileInfo, error)
	Update(*models.FileInfo) error
	Upsert(*models.FileInfo) error

	CreateMany([]*models.FileInfo) error
	UpsertMany([]*models.FileInfo) error
}

type IRepository interface {
	create(*models.FileInfo) error
	get(*string) (*models.FileInfo, error)
	update(*models.FileInfo) error
	upsert(*models.FileInfo) error

	createMany([]*models.FileInfo) error
	upsertMany([]*models.FileInfo) error
}
type IFilesService interface {
	GetFullPath(*string) *string
}

type Handler struct {
	service      IService
	filesService IFilesService
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
	return NewHandler(NewService(NewRepository(db)), NewFilesService(uploader))
}

// NewHandler constructor
func NewHandler(service IService, filesService IFilesService) *Handler {
	return &Handler{service: service, filesService: filesService}
}

func CreateService(db *bun.DB) *Service {
	repo := NewRepository(db)
	return NewService(repo)
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
