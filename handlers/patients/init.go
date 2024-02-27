package patients

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/handlers/basehandler"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/uptrace/bun"
)

type IHandler interface {
	basehandler.IHandler
	FTSP(c *gin.Context)
	GetBySnilsNumber(c *gin.Context)
	GetActualAnthropomethry(c *gin.Context)
}

type IService interface {
	basehandler.IServiceWithContext[models.Patient, models.Patients, models.PatientsWithCount]
	GetBySnilsNumber(c context.Context, snilsNumber string) (*models.Patient, bool, error)
	GetActualAnthropomethry(c context.Context, patientId string) (uint, uint, *time.Time, error)
}

type IRepository interface {
	basehandler.IRepositoryWithContext[models.Patient, models.Patients, models.PatientsWithCount]
	GetBySnilsNumber(c context.Context, snilsNumber string) (*models.Patient, error)
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
	err        error
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
	tx     *bun.Tx
	Error  error
}

type FilesService struct {
	helper *helper.Helper
}

var (
	H *Handler
	S *Service
	R *Repository
	F *FilesService
)

func Init(h *helper.Helper) {
	R = NewRepository(h)
	S = NewService(R, h)
	F = NewFilesService(h)
	H = NewHandler(S, F, h)
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
