package patientsresearches

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/handlers/basehandler"
	"github.com/pro-assistance/pro-assister/helper"
)

type IHandler interface {
	basehandler.IHandler
	GetPatientResearch(c *gin.Context)
}

type IService interface {
	basehandler.IService[models.PatientResearch, models.PatientsResearches, models.PatientsResearchesWithCount]
	GetPatientResearch(c context.Context, patientId string, researchId string) (*models.PatientResearch, error)
}

type IRepository interface {
	basehandler.IRepository[models.PatientResearch, models.PatientsResearches, models.PatientsResearchesWithCount]
	GetPatientResearch(c context.Context, patientId string, researchId string) (*models.PatientResearch, error)
}

type IFilesService interface {
	basehandler.IFilesService
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

func CreateService(helper *helper.Helper) *Service {
	repo := NewRepository(helper)
	return NewService(repo, helper)
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
