package xlsx

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/helpers/xlsxhelper"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	RegisterQuery(c *gin.Context)
}

type IXlsxService interface {
	GetFile() ([]byte, error)
}

type IService interface {
	GetRegisterQuery(id string) (*models.ResearchQuery, error)
}

type IRepository interface {
	db() *bun.DB
}

type Handler struct {
	service     IService
	xlsxService IXlsxService
	helper      *helper.Helper
}

type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

type ServiceXLSX struct {
	xlsxHelper xlsxhelper.IXlsxHelper
}

func CreateHandler(h xlsxhelper.IXlsxHelper) *Handler {
	repo := NewRepository()
	service := NewService(repo)
	xlsxService := NewXlsxService(h)
	return NewHandler(service, xlsxService)
}

func NewHandler(service IService, s IXlsxService) *Handler {
	return &Handler{service: service, xlsxService: s}
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func NewRepository() *Repository {
	return &Repository{ctx: context.Background()}
}

func NewXlsxService(h xlsxhelper.IXlsxHelper) *ServiceXLSX {
	return &ServiceXLSX{xlsxHelper: h}
}
