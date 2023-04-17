package routing

import (
	"mdgkb/tsr-tegister-server-v1/handlers/auth"
	"mdgkb/tsr-tegister-server-v1/handlers/chopscalequestions"
	"mdgkb/tsr-tegister-server-v1/handlers/documenttypes"
	"mdgkb/tsr-tegister-server-v1/handlers/drug"
	"mdgkb/tsr-tegister-server-v1/handlers/fileinfos"
	"mdgkb/tsr-tegister-server-v1/handlers/hmfsescalequestions"
	"mdgkb/tsr-tegister-server-v1/handlers/insurancecompany"
	"mdgkb/tsr-tegister-server-v1/handlers/meta"
	"mdgkb/tsr-tegister-server-v1/handlers/mkbitems"
	"mdgkb/tsr-tegister-server-v1/handlers/patients"
	"mdgkb/tsr-tegister-server-v1/handlers/questions"
	"mdgkb/tsr-tegister-server-v1/handlers/regions"
	"mdgkb/tsr-tegister-server-v1/handlers/register"
	"mdgkb/tsr-tegister-server-v1/handlers/registerpropertytouser"
	"mdgkb/tsr-tegister-server-v1/handlers/registerquery"
	"mdgkb/tsr-tegister-server-v1/handlers/registertopatient"
	"mdgkb/tsr-tegister-server-v1/handlers/representative"
	"mdgkb/tsr-tegister-server-v1/handlers/representativetypes"
	"mdgkb/tsr-tegister-server-v1/handlers/researches"
	"mdgkb/tsr-tegister-server-v1/handlers/researchespools"
	"mdgkb/tsr-tegister-server-v1/handlers/search"
	"mdgkb/tsr-tegister-server-v1/handlers/users"
	authRouter "mdgkb/tsr-tegister-server-v1/routing/auth"
	chopScaleQuestionsRouter "mdgkb/tsr-tegister-server-v1/routing/chopscalequestions"
	documentTypesRouter "mdgkb/tsr-tegister-server-v1/routing/documenttypes"
	drugRouter "mdgkb/tsr-tegister-server-v1/routing/drug"
	fileInfoRouter "mdgkb/tsr-tegister-server-v1/routing/fileinfo"
	hmfseScaleQuestionsRouter "mdgkb/tsr-tegister-server-v1/routing/hmfsescalequestions"
	insuranceCompanyRouter "mdgkb/tsr-tegister-server-v1/routing/insurancecompany"
	metaRouter "mdgkb/tsr-tegister-server-v1/routing/meta"
	mkbItemsRouter "mdgkb/tsr-tegister-server-v1/routing/mkbitems"
	patientsRouter "mdgkb/tsr-tegister-server-v1/routing/patients"
	regionsRouter "mdgkb/tsr-tegister-server-v1/routing/regions"
	registerRouter "mdgkb/tsr-tegister-server-v1/routing/register"
	registerGroupRouter "mdgkb/tsr-tegister-server-v1/routing/registergroup"
	registerPropertyRouter "mdgkb/tsr-tegister-server-v1/routing/registerproperty"
	registerPropertyToUserRouter "mdgkb/tsr-tegister-server-v1/routing/registerpropertytouser"
	registerQueryRouter "mdgkb/tsr-tegister-server-v1/routing/registerquery"
	registerToPatientRouter "mdgkb/tsr-tegister-server-v1/routing/registertopatient"
	representativeRouter "mdgkb/tsr-tegister-server-v1/routing/representative"
	representativeTypesRouter "mdgkb/tsr-tegister-server-v1/routing/representativetypes"
	researchesRouter "mdgkb/tsr-tegister-server-v1/routing/researches"
	researchesPoolsRouter "mdgkb/tsr-tegister-server-v1/routing/researchespools"
	searchRouter "mdgkb/tsr-tegister-server-v1/routing/search"
	usersRouter "mdgkb/tsr-tegister-server-v1/routing/users"

	"github.com/gin-gonic/gin"

	helperPack "github.com/pro-assistance/pro-assister/helper"
)

func Init(r *gin.Engine, helper *helperPack.Helper) {
	r.Static("/api/v1/static", "./static/")
	//r.Use(helper.HTTP.CORSMiddleware())
	api := r.Group("/api/v1")
	authRouter.Init(api.Group("/auth"), auth.CreateHandler(helper))
	documentTypesRouter.Init(api.Group("/document-types"), documenttypes.CreateHandler(helper))
	drugRouter.Init(api.Group("/drugs"), drug.CreateHandler(helper))
	fileInfoRouter.Init(api.Group("/files-info"), fileinfos.CreateHandler(helper))
	insuranceCompanyRouter.Init(api.Group("/insurance-companies"), insurancecompany.CreateHandler(helper))
	metaRouter.Init(api.Group("/meta"), meta.CreateHandler(helper))
	mkbItemsRouter.Init(api.Group("/mkb-items"), mkbitems.CreateHandler(helper))
	patientsRouter.Init(api.Group("/patients"), patients.CreateHandler(helper))
	registerRouter.Init(api.Group("/registers"), register.CreateHandler(helper))
	researchesRouter.Init(api.Group("/researches"), researches.CreateHandler(helper))
	registerToPatientRouter.Init(api.Group("/registers-to-patients"), registertopatient.CreateHandler(helper))
	registerGroupRouter.Init(api.Group("/register-groups"), questions.CreateHandler(helper))
	registerQueryRouter.Init(api.Group("/register-queries"), registerquery.CreateHandler(helper))
	registerPropertyRouter.Init(api.Group("/register-properties"), questions.CreateHandler(helper))
	representativeRouter.Init(api.Group("/representatives"), representative.CreateHandler(helper))
	representativeTypesRouter.Init(api.Group("/representative-types"), representativetypes.CreateHandler(helper))
	registerPropertyToUserRouter.Init(api.Group("/register-properties-to-user"), registerpropertytouser.CreateHandler(helper))
	usersRouter.Init(api.Group("/users"), users.CreateHandler(helper))
	regionsRouter.Init(api.Group("/regions"), regions.CreateHandler(helper))
	researchesPoolsRouter.Init(api.Group("/researches-pools"), researchespools.CreateHandler(helper))
	searchRouter.Init(api.Group("/search"), search.CreateHandler(helper))
	chopScaleQuestionsRouter.Init(api.Group("/chop-scale-questions"), chopscalequestions.CreateHandler(helper))
	hmfseScaleQuestionsRouter.Init(api.Group("/hmfse-scale-questions"), hmfsescalequestions.CreateHandler(helper))
}
