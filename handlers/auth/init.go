package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

type IHandler interface {
	Register(c *gin.Context) error
	Login(c *gin.Context) error
	//Refresh(c *gin.Context) error
	//Logout(c *gin.Context) error
}

type IService interface {
	Register(*models.User) (*models.TokensWithUser, error)
	Login(*models.User) (*models.TokensWithUser, error)
}

type IRepository interface {
	getByLogin(*string) (*models.User, error)
	create(*models.User) error
}

type Handler struct {
	service IService
}

type Service struct {
	repository IRepository
	redis      *redis.Client
}

type Repository struct {
	db  *bun.DB
	ctx context.Context
}

func CreateHandler(db *bun.DB, redisClient *redis.Client) *Handler {
	repo := NewRepository(db)
	service := NewService(repo, redisClient)
	return NewHandler(service)
}

// NewHandler constructor
func NewHandler(s IService) *Handler {
	return &Handler{service: s}
}

func NewService(repository IRepository, redisClient *redis.Client) *Service {
	return &Service{repository: repository, redis: redisClient}
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, ctx: context.Background()}
}
