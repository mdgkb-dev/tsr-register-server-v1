package routing

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	helperPack "github.com/pro-assistance/pro-assister/helper"
	"mdgkb/tsr-tegister-server-v1/handlers/auth"
	"mdgkb/tsr-tegister-server-v1/handlers/documentTypes"
	"mdgkb/tsr-tegister-server-v1/handlers/drug"
	"mdgkb/tsr-tegister-server-v1/handlers/fileInfo"
	"mdgkb/tsr-tegister-server-v1/handlers/insuranceCompany"
	"mdgkb/tsr-tegister-server-v1/handlers/meta"
	"mdgkb/tsr-tegister-server-v1/handlers/mkb"
	"mdgkb/tsr-tegister-server-v1/handlers/mkbConcreteDiagnoses"
	"mdgkb/tsr-tegister-server-v1/handlers/mkbDiagnoses"
	"mdgkb/tsr-tegister-server-v1/handlers/mkbGroups"
	"mdgkb/tsr-tegister-server-v1/handlers/mkbSubDiagnoses"
	"mdgkb/tsr-tegister-server-v1/handlers/patients"
	"mdgkb/tsr-tegister-server-v1/handlers/regions"
	"mdgkb/tsr-tegister-server-v1/handlers/register"
	"mdgkb/tsr-tegister-server-v1/handlers/registerProperty"
	"mdgkb/tsr-tegister-server-v1/handlers/registerPropertyToUser"
	"mdgkb/tsr-tegister-server-v1/handlers/registerQuery"
	"mdgkb/tsr-tegister-server-v1/handlers/representative"
	"mdgkb/tsr-tegister-server-v1/handlers/representativeTypes"
	"mdgkb/tsr-tegister-server-v1/handlers/search"
	"mdgkb/tsr-tegister-server-v1/handlers/users"
	"mdgkb/tsr-tegister-server-v1/middleware"
	authRouter "mdgkb/tsr-tegister-server-v1/routing/auth"
	documentTypesRouter "mdgkb/tsr-tegister-server-v1/routing/documentTypes"
	drugRouter "mdgkb/tsr-tegister-server-v1/routing/drug"
	fileInfoRouter "mdgkb/tsr-tegister-server-v1/routing/fileInfo"
	insuranceCompanyRouter "mdgkb/tsr-tegister-server-v1/routing/insuranceCompany"
	metaRouter "mdgkb/tsr-tegister-server-v1/routing/meta"
	mkbRouter "mdgkb/tsr-tegister-server-v1/routing/mkb"
	mkbConcreteDiagnosesRouter "mdgkb/tsr-tegister-server-v1/routing/mkbConcreteDiagnoses"
	mkbDiagnosesRouter "mdgkb/tsr-tegister-server-v1/routing/mkbDiagnoses"
	mkbGroupsRouter "mdgkb/tsr-tegister-server-v1/routing/mkbGroups"
	mkbSubDiagnosesRouter "mdgkb/tsr-tegister-server-v1/routing/mkbSubDiagnoses"
	patientsRouter "mdgkb/tsr-tegister-server-v1/routing/patients"
	regionsRouter "mdgkb/tsr-tegister-server-v1/routing/regions"
	registerRouter "mdgkb/tsr-tegister-server-v1/routing/register"
	registerGroupRouter "mdgkb/tsr-tegister-server-v1/routing/registerGroup"
	registerPropertyRouter "mdgkb/tsr-tegister-server-v1/routing/registerProperty"
	registerPropertyToUserRouter "mdgkb/tsr-tegister-server-v1/routing/registerPropertyToUser"
	registerQueryRouter "mdgkb/tsr-tegister-server-v1/routing/registerQuery"
	representativeRouter "mdgkb/tsr-tegister-server-v1/routing/representative"
	representativeTypesRouter "mdgkb/tsr-tegister-server-v1/routing/representativeTypes"
	searchRouter "mdgkb/tsr-tegister-server-v1/routing/search"
	usersRouter "mdgkb/tsr-tegister-server-v1/routing/users"
)

func Init(r *gin.Engine, helper *helperPack.Helper) {
	r.Static("/static", "../static/")
	api := r.Group("/api/v1")
	m := middleware.CreateMiddleware(helper)
	r.Use(m.CORSMiddleware())
	authRouter.Init(api.Group("/auth"), auth.CreateHandler(helper))
	documentTypesRouter.Init(api.Group("/document-types"), documentTypes.CreateHandler(helper))
	drugRouter.Init(api.Group("/drugs"), drug.CreateHandler(helper))
	fileInfoRouter.Init(api.Group("/files-info"), fileInfo.CreateHandler(helper))
	insuranceCompanyRouter.Init(api.Group("/insurance-companies"), insuranceCompany.CreateHandler(helper))
	metaRouter.Init(api.Group("/meta"), meta.CreateHandler(helper))
	mkbRouter.Init(api.Group("/mkb"), mkb.CreateHandler(helper))
	patientsRouter.Init(api.Group("/patients"), patients.CreateHandler(helper))
	registerRouter.Init(api.Group("/registers"), register.CreateHandler(helper))
	registerGroupRouter.Init(api.Group("/register-groups"), registerProperty.CreateHandler(helper))
	registerQueryRouter.Init(api.Group("/register-queries"), registerQuery.CreateHandler(helper))
	registerPropertyRouter.Init(api.Group("/register-properties"), registerProperty.CreateHandler(helper))
	representativeRouter.Init(api.Group("/representatives"), representative.CreateHandler(helper))
	representativeTypesRouter.Init(api.Group("/representative-types"), representativeTypes.CreateHandler(helper))
	registerPropertyToUserRouter.Init(api.Group("/register-properties-to-user"), registerPropertyToUser.CreateHandler(helper))
	usersRouter.Init(api.Group("/users"), users.CreateHandler(helper))
	regionsRouter.Init(api.Group("/regions"), regions.CreateHandler(helper))
	searchRouter.Init(api.Group("/search"), search.CreateHandler(helper))
	mkbGroupsRouter.Init(api.Group("/mkb-groups"), mkbGroups.CreateHandler(helper))
	mkbDiagnosesRouter.Init(api.Group("/mkb-diagnoses"), mkbDiagnoses.CreateHandler(helper))
	mkbSubDiagnosesRouter.Init(api.Group("/mkb-sub-diagnoses"), mkbSubDiagnoses.CreateHandler(helper))
	mkbConcreteDiagnosesRouter.Init(api.Group("/mkb-concrete-diagnoses"), mkbConcreteDiagnoses.CreateHandler(helper))
}
