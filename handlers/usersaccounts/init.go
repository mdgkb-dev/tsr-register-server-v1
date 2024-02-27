package usersaccounts

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/uptrace/bun"
)

type IService interface {
	SetQueryFilter(*gin.Context) error
	Create(user *models.UserAccount) error
	CheckAccountPassword(user *models.UserAccount, skipPassword bool) error
	GetByEmail(email string) (*models.UserAccount, error)
	Get(id string) (*models.UserAccount, error)
	UpdateUUID(id string) error
	UpdatePassword(account *models.UserAccount) error
}

type IRepository interface {
	SetQueryFilter(*gin.Context) error
	DB() *bun.DB
	Create(account *models.UserAccount) error
	GetByEmail(string) (*models.UserAccount, error)
	Get(string) (*models.UserAccount, error)
	UpdateUUID(id string) error
	UpdatePassword(account *models.UserAccount) error
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
	ctx    context.Context
	helper *helper.Helper
}

type FilesService struct {
	helper *helper.Helper
}

type ValidateService struct {
	helper *helper.Helper
}

func CreateService(helper *helper.Helper) *Service {
	return NewService(NewRepository(helper), helper)
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
