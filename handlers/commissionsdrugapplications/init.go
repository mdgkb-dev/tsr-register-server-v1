package commissionsdrugapplications

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"
	"mime/multipart"

	"github.com/pro-assistance/pro-assister/handlers/basehandler"
	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	basehandler.IHandler
}

type IService interface {
	basehandler.IService[models.CommissionDrugApplication, models.CommissionsDrugApplications, models.CommissionsDrugApplicationsWithCount]
}

type IRepository interface {
	basehandler.IRepository[models.CommissionDrugApplication, models.CommissionsDrugApplications, models.CommissionsDrugApplicationsWithCount]
}

type IFilesService interface {
	Upload(*gin.Context, *models.CommissionDrugApplication, map[string][]*multipart.FileHeader) error
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
