package mkb

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAllClasses(c *gin.Context)
	GetGroupByClassId(c *gin.Context)
	GetGroupChildrens(c *gin.Context)
	GetSubGroupChildrens(c *gin.Context)
	GetGroupsBySearch(c *gin.Context)
	GetDiagnosisBySearch(c *gin.Context)
	GetDiagnosisByGroupId(c *gin.Context)
	GetSubDiagnosisByDiagnosisId(c *gin.Context)
}

type IService interface {
	GetAllClasses() ([]*models.MkbClass, error)
	GetGroupByClassId(*string) (*CompositionMkb, error)
	GetGroupChildrens(*string) (*CompositionMkb, error)
	GetDiagnosisByGroupId(*string) ([]*models.MkbDiagnosis, error)
	GetSubGroupChildrens(*string) (*CompositionMkb, error)
	GetSubDiagnosisByDiagnosisId(*string) ([]*models.MkbSubDiagnosis, error)
	GetDiagnosisBySearch(*string) ([]*models.MkbDiagnosis, error)
	GetGroupsBySearch(*string) ([]*models.MkbGroup, error)
	UpdateRelevant(*string, *string) error
	UpdateName(*string, *string, *string) error
}

type IRepository interface {
	getDB() *bun.DB
	getAllClasses() (items []*models.MkbClass, err error)
	getGroupsByClassId(*string) (items []*models.MkbGroup, err error)
	getSubGroupByGroupId(*string) (items []*models.MkbSubGroup, err error)
	getDiagnosisByClassId(*string) (items []*models.MkbDiagnosis, err error)
	getDiagnosisByGroupId(*string) (items []*models.MkbDiagnosis, err error)
	getDiagnosisBySubGroupId(*string) (items []*models.MkbDiagnosis, err error)
	getDiagnosisBySubSubGroupId(*string) (items []*models.MkbDiagnosis, err error)
	getSubDiagnosisByDiagnosisId(*string) (items []*models.MkbSubDiagnosis, err error)
	getGroupsByRange(*string) (items []*models.MkbGroup, err error)
	getGroupBySearch(*string) (items []*models.MkbGroup, err error)
	getDiagnosisBySearch(*string) (items []*models.MkbDiagnosis, err error)
	updateRelevant(*string, *string) error
	updateName(*string, *string, *string) error
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
	db     *bun.DB
	ctx    context.Context
	helper *helper.Helper
}

func CreateHandler(db *bun.DB, helper *helper.Helper) *Handler {
	repo := NewRepository(db, helper)
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

func NewRepository(db *bun.DB, helper *helper.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}
