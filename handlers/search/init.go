package search

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/search"
	"mime/multipart"

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
	SearchObjects(*search.SearchModel) error
	SearchGroups() (search.SearchGroups, error)
	Search(*search.SearchModel) error
}

type IRepository interface {
	getDB() *bun.DB
	getGroups(string) (search.SearchGroups, error)
	search(*search.SearchModel) error
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
	db            *bun.DB
	ctx           context.Context
	helper        *helper.Helper
	elasticsearch *elasticsearch.Client
}

type FilesService struct {
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

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}
