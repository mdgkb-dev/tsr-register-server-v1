package patientsdomains

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/httpHelper/basehandler"
	"github.com/pro-assistance/pro-assister/sqlHelper"
	"github.com/uptrace/bun"
)

type IHandler interface {
	basehandler.IHandler
	AddToDomain(c *gin.Context)
}

type IService interface {
	basehandler.IServiceWithContext[models.PatientDomain, models.PatientsDomains, models.PatientsDomainsWithCount]
	PatientInDomain(c context.Context, patientID string) (bool, error)
}

type IRepository interface {
	basehandler.IRepositoryWithContext[models.PatientDomain, models.PatientsDomains, models.PatientsDomainsWithCount]
	PatientInDomain(c context.Context, patientID string) (bool, error)
}

type IFilesService interface {
	basehandler.IFilesService
}

type Handler struct {
	service      IService
	filesService IFilesService
	helper       *helper.Helper
}

var H *Handler
var S *Service
var R *Repository
var F *FilesService

type Service struct {
	Repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx         context.Context
	helper      *helper.Helper
	queryFilter *sqlHelper.QueryFilter
	Tx          *bun.Tx
}

type FilesService struct {
	helper *helper.Helper
}

func Init(h *helper.Helper) {
	R = NewRepository(h)
	S = NewService(R, h)
	F = NewFilesService(h)
	H = NewHandler(S, F, h)
}

// NewHandler constructor
func NewHandler(s IService, filesService IFilesService, helper *helper.Helper) *Handler {
	return &Handler{service: s, filesService: filesService, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{Repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
