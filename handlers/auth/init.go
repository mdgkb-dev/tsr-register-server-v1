package auth

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"
	"github.com/uptrace/bun"
)

type IHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	LoginAs(c *gin.Context)
	Logout(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type IService interface {
	setQueryFilter(*gin.Context) error
	Register(user *models.UserAccount) (*models.TokensWithUser, error)
	Login(user *models.UserAccount, skipPassword bool) (*models.TokensWithUser, error)
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	db() *bun.DB
}

type IValidator interface {
	Login(*models.UserAccount) error
}

type Handler struct {
	service   IService
	helper    *helper.Helper
	validator IValidator
}

type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	ctx         context.Context
	helper      *helper.Helper
	queryFilter *sqlHelper.QueryFilter
}

type FilesService struct {
	helper *helper.Helper
}

type ValidateService struct {
	helper *helper.Helper
}

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
	service := NewService(repo, helper)
	return NewHandler(service, helper)
}

// NewHandler constructor
func NewHandler(s IService, helper *helper.Helper) *Handler {
	return &Handler{service: s, helper: helper}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(helper *helper.Helper) *Repository {
	return &Repository{ctx: context.Background(), helper: helper}
}
