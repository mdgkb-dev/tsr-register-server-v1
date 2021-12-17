package users

import (
	"context"
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
}

type IService interface {
	GetAll() (models.Users, error)
	Get(*string) (*models.User, error)
	Create(*models.User) error
	Update(*models.User) error
	Delete(*string) error
}

type IRepository interface {
	getDB() *bun.DB
	create(*models.User) error
	getAll() (models.Users, error)
	get(*string) (*models.User, error)
	update(*models.User) error
	delete(*string) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.User, map[string][]*multipart.FileHeader) error
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
	repo := NewRepository(db)
	service := NewService(repo)
	filesService := NewFilesService(uploader)
	return NewHandler(service, filesService)
}

// NewHandler constructor
func NewHandler(service IService, filesService IFilesService) *Handler {
	return &Handler{service: service, filesService: filesService}
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
