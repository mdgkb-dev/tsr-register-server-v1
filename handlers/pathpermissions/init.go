package pathpermissions

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"
	"github.com/uptrace/bun"
)

type IHandler interface {
	SavePathPermissions(c *gin.Context)
	GetAllPathPermissions(c *gin.Context)
	GetAllPathPermissionsAdmin(c *gin.Context)
	GetPathPermissionsByRoleID(c *gin.Context)
	CheckPathPermissions(c *gin.Context)
}

type IService interface {
	setQueryFilter(*gin.Context) error

	UpsertManyPathPermissions(models.PathPermissions) error
	GetAllPathPermissions() (models.PathPermissions, error)
	GetAllPathPermissionsAdmin() (models.PathPermissionsWithCount, error)
	GetPathPermissionsByRoleID(id string) (models.PathPermissions, error)
	CheckPathPermissions(path string, roleID string) error
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	db() *bun.DB
	getAllPathPermissions() (models.PathPermissions, error)
	getAllPathPermissionsAdmin() (models.PathPermissionsWithCount, error)
	getPathPermissionsByRoleID(id string) (models.PathPermissions, error)
	upsertManyPathPermissions(items models.PathPermissions) (err error)
	deleteManyPathPermissions(idPool []uuid.UUID) (err error)
	upsertManyPathPermissionsRoles(items models.PathPermissionsRoles) (err error)
	deleteManyPathPermissionsRoles(idPool []uuid.UUID) (err error)
	checkPathPermissions(path string, roleID string) error
}

type Handler struct {
	service IService
	helper  *helper.Helper
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
