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
	Update(c *gin.Context)
	GetGroupByClassId(c *gin.Context)
	GetGroupChildrens(c *gin.Context)
	GetSubGroupChildrens(c *gin.Context)
	GetGroupsBySearch(c *gin.Context)
	GetDiagnosisBySearch(c *gin.Context)
	GetSubDiagnosesBySearch(c *gin.Context)
	GetConcreteDiagnosisBySearch(c *gin.Context)
	GetDiagnosisByGroupId(c *gin.Context)
	GetSubDiagnosisByDiagnosisId(c *gin.Context)
	GetConcreteDiagnosisBySubDiagnosisId(c *gin.Context)
}

type IService interface {
	GetAllClasses() (models.MkbClasses, error)
	GetGroupByClassId(string) (*CompositionMkb, error)
	GetGroupChildrens(string) (*CompositionMkb, error)
	GetDiagnosisByGroupId(string) (models.MkbDiagnoses, error)
	GetSubGroupChildrens(string) (*CompositionMkb, error)
	GetSubDiagnosisByDiagnosisId(string) (models.MkbSubDiagnoses, error)
	GetDiagnosisBySearch(string) (models.MkbDiagnoses, error)
	GetGroupsBySearch(string) (models.MkbGroups, error)
	UpdateRelevant(string, string) error
	UpdateName(string, string, string) error

	GetSubDiagnosesBySearch(string) (models.MkbSubDiagnoses, error)
	GetConcreteDiagnosisBySearch(string) (models.MkbConcreteDiagnoses, error)
	GetConcreteDiagnosisBySubDiagnosisId(string) (models.MkbConcreteDiagnoses, error)
}

type IRepository interface {
	getDB() *bun.DB
	getAllClasses() (items models.MkbClasses, err error)
	getGroupsByClassId(string) (items models.MkbGroups, err error)
	getSubGroupByGroupId(string) (items models.MkbSubGroups, err error)
	getDiagnosisByClassId(string) (items models.MkbDiagnoses, err error)
	getDiagnosisByGroupId(string) (items models.MkbDiagnoses, err error)
	getDiagnosisBySubGroupId(string) (items models.MkbDiagnoses, err error)
	getDiagnosisBySubSubGroupId(string) (items models.MkbDiagnoses, err error)
	getSubDiagnosisByDiagnosisId(string) (items models.MkbSubDiagnoses, err error)

	getConcreteDiagnosisBySearch(string) (models.MkbConcreteDiagnoses, error)
	getSubDiagnosesBySearch(string) (models.MkbSubDiagnoses, error)
	getConcreteDiagnosisBySubDiagnosisId(string) (models.MkbConcreteDiagnoses, error)

	getGroupsByRange(string) (items models.MkbGroups, err error)
	getGroupBySearch(string) (items models.MkbGroups, err error)
	getDiagnosisBySearch(string) (items models.MkbDiagnoses, err error)
	updateRelevant(string, string) error
	updateName(string, string, string) error
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
