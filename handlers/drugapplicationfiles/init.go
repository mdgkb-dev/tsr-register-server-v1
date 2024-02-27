package drugapplicationfiles

import (
	"context"

	"github.com/pro-assistance/pro-assister/handlers/basehandler"
	"github.com/pro-assistance/pro-assister/helper"

	"mdgkb/tsr-tegister-server-v1/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	basehandler.IHandler
}

type IService interface {
	basehandler.IService[models.DrugApplicationFile, models.DrugApplicationFiles, models.DrugApplicationFilesWithCount]
}

type IRepository interface {
	basehandler.IService[models.DrugApplicationFile, models.DrugApplicationFiles, models.DrugApplicationFilesWithCount]
}

type IFilesService interface {
	Upload(*gin.Context, *models.DrugApplicationFile, map[string][]*multipart.FileHeader) error
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
