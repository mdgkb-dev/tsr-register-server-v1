package routing

import (
	"mdgkb/tsr-tegister-server-v1/config"
	"mdgkb/tsr-tegister-server-v1/helpers/uploadHelper"
	"mdgkb/tsr-tegister-server-v1/routing/auth"
	"mdgkb/tsr-tegister-server-v1/routing/documentTypes"
	"mdgkb/tsr-tegister-server-v1/routing/drug"
	"mdgkb/tsr-tegister-server-v1/routing/fileInfo"
	"mdgkb/tsr-tegister-server-v1/routing/insuranceCompany"
	"mdgkb/tsr-tegister-server-v1/routing/meta"
	"mdgkb/tsr-tegister-server-v1/routing/mkb"
	"mdgkb/tsr-tegister-server-v1/routing/patient"
	"mdgkb/tsr-tegister-server-v1/routing/register"
	"mdgkb/tsr-tegister-server-v1/routing/registerGroup"
	"mdgkb/tsr-tegister-server-v1/routing/registerProperty"
	"mdgkb/tsr-tegister-server-v1/routing/representative"
	"mdgkb/tsr-tegister-server-v1/routing/representativeTypes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/go-redis/redis/v7"
	"github.com/uptrace/bun"
)

func Init(r *gin.Engine, db *bun.DB, redisClient *redis.Client, config config.Config) {
	localUploader := uploadHelper.NewLocalUploader(&config.UploadPath)
	r.Static("/static", "./static/")
	api := r.Group("/api/v1")

	auth.Init(api.Group("/auth"), db, redisClient)
	documentTypes.Init(api.Group("/document-types"), db)
	drug.Init(api.Group("/drugs"), db)
	fileInfo.Init(api.Group("/files-info"), db, localUploader)
	insuranceCompany.Init(api.Group("/insurance-companies"), db)
	meta.Init(api.Group("/meta"), db)
	mkb.Init(api.Group("/mkb"), db)
	patient.Init(api.Group("/patients"), db, localUploader)
	register.Init(api.Group("/registers"), db)
	registerGroup.Init(api.Group("/register-groups"), db)
	registerProperty.Init(api.Group("/register-properties"), db)
	representative.Init(api.Group("/representatives"), db, localUploader)
	representativeTypes.Init(api.Group("/representative-types"), db)
}
