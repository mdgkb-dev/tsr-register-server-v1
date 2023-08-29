package questions

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/httpHelper/basehandler"
	"github.com/pro-assistance/pro-assister/sqlHelper"
)

type IHandler interface {
	basehandler.IHandler
}

type IService interface {
	basehandler.IServiceWithContext[models.Question, models.Questions, models.QuestionsWithCount]

	UpsertMany(context.Context, models.Questions) error
	DeleteMany(context.Context, []uuid.UUID) error
}

type IRepository interface {
	basehandler.IRepositoryWithContext[models.Question, models.Questions, models.QuestionsWithCount]

	upsertMany(context.Context, models.Questions) error
	deleteMany(context.Context, []uuid.UUID) error
}

var H *Handler
var S *Service
var R *Repository

func Init(h *helper.Helper) {
	R = NewRepository(h)
	S = NewService(R, h)
	H = NewHandler(S, h)
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

func CreateHandler(helper *helper.Helper) *Handler {
	repo := NewRepository(helper)
	service := NewService(repo, helper)
	return NewHandler(service, helper)
}

func CreateService(helper *helper.Helper) *Service {
	repo := NewRepository(helper)
	return NewService(repo, helper)
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
