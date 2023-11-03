package search

import (
	"context"
	"mime/multipart"

	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/search"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	Search(c *gin.Context)
	ElasticSearch(c *gin.Context)
	SearchMain(c *gin.Context)
	SearchGroups(c *gin.Context)
}

type IService interface {
	SearchMain(*search.SearchModel) error
	SearchObjects(context.Context, *search.SearchModel) error
	SearchGroups() (search.SearchGroups, error)
	Search(*search.SearchModel) error
}

type IRepository interface {
	db() *bun.DB
	getGroups(string) (search.SearchGroups, error)
	search(context.Context,*search.SearchModel) error
	elasticSearch(*search.SearchModel) error
	elasticSuggester(*search.SearchModel) error
}

type IFilesService interface {
	Upload(*gin.Context, interface{}, map[string][]*multipart.FileHeader) error
}

type Handler struct {
	service      IService
	filesService IFilesService
	helper       *helper.Helper
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

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
