package xlsx

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/tsr-tegister-server-v1/helpers/xlsxHelper"
	"mdgkb/tsr-tegister-server-v1/models"

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
	GetRegisterQuery(id string) (*models.RegisterQuery, error)
}

type IRepository interface {
	getDB() *bun.DB
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
	db     *bun.DB
	ctx    context.Context
	helper *helper.Helper
}

type XlsxService struct {
	xlsxHelper xlsxHelper.IXlsxHelper
}

func CreateHandler(db *bun.DB, h xlsxHelper.IXlsxHelper) *Handler {
	repo := NewRepository(db)
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

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, ctx: context.Background()}
}

func NewXlsxService(h xlsxHelper.IXlsxHelper) *XlsxService {
	return &XlsxService{xlsxHelper: h}
}
