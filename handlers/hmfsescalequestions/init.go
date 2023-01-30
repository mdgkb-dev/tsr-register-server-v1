package hmfsescalequestions

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"
	"mime/multipart"

	"github.com/pro-assistance/pro-assister/helper"

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
	GetAll() (models.HmfseScaleQuestions, error)
	Get(string) (*models.HmfseScaleQuestion, error)
	Create(*models.HmfseScaleQuestion) error
	Update(item *models.HmfseScaleQuestion) error
	Delete(string) error
}

type IRepository interface {
	db() *bun.DB
	create(*models.HmfseScaleQuestion) error
	getAll() (models.HmfseScaleQuestions, error)
	get(string) (*models.HmfseScaleQuestion, error)
	update(item *models.HmfseScaleQuestion) error
	delete(string) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.HmfseScaleQuestion, map[string][]*multipart.FileHeader) error
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

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
