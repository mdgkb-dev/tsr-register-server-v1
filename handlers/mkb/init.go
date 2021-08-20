package mkb

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"mdgkb/tsr-tegister-server-v1/models"
)

type IHandler interface {
	GetAllClasses(c *gin.Context) error
	GetGroupByClassId(c *gin.Context) error
	GetGroupChildrens(c *gin.Context) error
	GetSubGroupChildrens(c *gin.Context) error
	GetGroupsBySearch(c *gin.Context) error
	GetDiagnosisBySearch(c *gin.Context) error
	GetDiagnosisByGroupId(c *gin.Context) error
	GetSubDiagnosisByDiagnosisId(c *gin.Context) error
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
}

type Service struct {
	repository IRepository
}

type Repository struct {
	db  *bun.DB
	ctx context.Context
}

func CreateHandler(db *bun.DB) *Handler {
	repo := NewRepository(db)
	service := NewService(repo)
	return NewHandler(service)
}

// NewHandler constructor
func NewHandler(s IService) *Handler {
	return &Handler{service: s}
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db: db, ctx: context.Background()}
}
