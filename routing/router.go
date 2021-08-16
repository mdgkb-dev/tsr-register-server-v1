package routing

import (
	"mdgkb/tsr-tegister-server-v1/config"
	"mdgkb/tsr-tegister-server-v1/helpers"
	"mdgkb/tsr-tegister-server-v1/routing/test"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/go-redis/redis/v7"
	"github.com/uptrace/bun"
)

func Init(r *gin.Engine, db *bun.DB, redisClient *redis.Client, config config.Config) {
	localUploader := helpers.NewLocalUploader(&config.UploadPath)
	r.Static("/static", "./static/")
	api := r.Group("/api/v1")

	test.Init(api.Group("/users"), db, localUploader)
}
