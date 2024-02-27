package users

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
	basehandler.IServiceWithContext[models.User, models.Users, models.UsersWithCount]
}

type IRepository interface {
	basehandler.IRepositoryWithContext[models.User, models.Users, models.UsersWithCount]
}

type IFilesService interface {
	Upload(*gin.Context, *models.User, map[string][]*multipart.FileHeader) error
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

var (
	H *Handler
	S *Service
	R *Repository
	F *FilesService
)

func Init(h *helper.Helper) {
	H = &Handler{helper: h}
	S = &Service{helper: h}
	R = &Repository{ctx: context.Background(), helper: h}
	F = &FilesService{helper: h}
}
