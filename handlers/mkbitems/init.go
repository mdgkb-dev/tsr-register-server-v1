package mkbitems

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/pro-assistance/pro-assister/helper"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	Get(c *gin.Context)
	//GetAllClasses(c *gin.Context)
	//Update(c *gin.Context)
	//GetGroupByClassID(c *gin.Context)
	//GetGroupChildrens(c *gin.Context)
	//GetSubGroupChildrens(c *gin.Context)
	//GetGroupsBySearch(c *gin.Context)
	//GetDiagnosisBySearch(c *gin.Context)
	//GetSubDiagnosesBySearch(c *gin.Context)
	//GetConcreteDiagnosisBySearch(c *gin.Context)
	//GetDiagnosisByGroupID(c *gin.Context)
	//GetSubDiagnosisByDiagnosisID(c *gin.Context)
	//GetConcreteDiagnosisBySubDiagnosisID(c *gin.Context)
	//SelectMkbElement(c *gin.Context)
	GetTree(c *gin.Context)
}

type IService interface {
	//GetAllClasses() (models.MkbClasses, error)
	//GetGroupByClassID(string) (*CompositionMkb, error)
	//GetGroupChildrens(string) (*CompositionMkb, error)
	//GetDiagnosisByGroupID(string) (models.MkbDiagnoses, error)
	//GetSubGroupChildrens(string) (*CompositionMkb, error)
	//GetSubDiagnosisByDiagnosisID(string) (models.MkbSubDiagnoses, error)
	//GetDiagnosisBySearch(string) (models.MkbDiagnoses, error)
	//GetGroupsBySearch(string) (models.MkbGroups, error)
	//UpdateRelevant(string, string) error
	//UpdateName(string, string, string) error
	//
	//GetSubDiagnosesBySearch(string) (models.MkbSubDiagnoses, error)
	//GetConcreteDiagnosisBySearch(string) (models.MkbConcreteDiagnoses, error)
	//GetConcreteDiagnosisBySubDiagnosisID(string) (models.MkbConcreteDiagnoses, error)
	//SelectMkbElement(string) (*models.MkbClass, *models.MkbElement, error)
	GetTree() (*models.MkbItem, error)
	Get(string) (*models.MkbItem, error)
}

type IRepository interface {
	db() *bun.DB
	//getAllClasses() (items models.MkbClasses, err error)
	//getGroupsByClassID(string) (items models.MkbGroups, err error)
	//getSubGroupByGroupID(string) (items models.MkbSubGroups, err error)
	//getDiagnosisByClassID(string) (items models.MkbDiagnoses, err error)
	//getDiagnosisByGroupID(string) (items models.MkbDiagnoses, err error)
	//getDiagnosisBySubGroupID(string) (items models.MkbDiagnoses, err error)
	//getDiagnosisBySubSubGroupID(string) (items models.MkbDiagnoses, err error)
	//getSubDiagnosisByDiagnosisID(string) (items models.MkbSubDiagnoses, err error)
	//
	//getConcreteDiagnosisBySearch(string) (models.MkbConcreteDiagnoses, error)
	//getSubDiagnosesBySearch(string) (models.MkbSubDiagnoses, error)
	//getConcreteDiagnosisBySubDiagnosisID(string) (models.MkbConcreteDiagnoses, error)
	//
	//getGroupsByRange(string) (items models.MkbGroups, err error)
	//getGroupBySearch(string) (items models.MkbGroups, err error)
	//getDiagnosisBySearch(string) (items models.MkbDiagnoses, err error)
	//updateRelevant(string, string) error
	//updateName(string, string, string) error
	//selectMkbElement(string) (*models.MkbElement, error)
	//getMkbClass(id uuid.UUID) (item *models.MkbClass, err error)
	GetTree() (models.MkbItems, error)
	Get(string) (*models.MkbItem, error)
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
	ctx    context.Context
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
