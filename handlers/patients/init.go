package patients

import (
	"context"

	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/httpHelper/basehandler"
	"github.com/pro-assistance/pro-assister/sqlHelper"
	"github.com/pro-assistance/pro-assister/tokenHelper"

	"mdgkb/tsr-tegister-server-v1/models"
)

type IHandler interface {
	basehandler.IHandler
}

type IService interface {
	basehandler.IService[models.Patient, models.Patients, models.PatientsWithCount]
}

type IRepository interface {
	basehandler.IRepository[models.Patient, models.Patients, models.PatientsWithCount]
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
	ctx           context.Context
	helper        *helper.Helper
	queryFilter   *sqlHelper.QueryFilter
	accessDetails *tokenHelper.AccessDetails
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
