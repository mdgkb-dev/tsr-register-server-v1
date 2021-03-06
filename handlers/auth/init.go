package auth

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

type IHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
	Me(c *gin.Context)
	DoesLoginExist(c *gin.Context)
	CheckPathPermissions(c *gin.Context)
	//Refresh(c *gin.Context) error
	//Logout(c *gin.Context) error
}

type IService interface {
	Register(*models.User) (*models.TokensWithUser, error)
	Login(*models.User) (*models.TokensWithUser, error)
	GetUserByID(*string) (*models.User, error)
	DoesLoginExist(*string) (bool, error)
}

type IRepository interface {
	getByLogin(*string) (*models.User, error)
	getByID(*string) (*models.User, error)
	create(*models.User) error
}

type Handler struct {
	service IService
	helper  *helper.Helper
}

type Service struct {
	repository IRepository
	redis      *redis.Client
	helper     *helper.Helper
}

type Repository struct {
	ctx    context.Context
	helper *helper.Helper
}

type DoesLoginExist struct {
	DoesLoginExist bool
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
