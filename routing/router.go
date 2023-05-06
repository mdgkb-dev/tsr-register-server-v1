package routing

import (
	"mdgkb/tsr-tegister-server-v1/handlers/anamneses"
	"mdgkb/tsr-tegister-server-v1/handlers/auth"
	"mdgkb/tsr-tegister-server-v1/handlers/chopscalequestions"
	"mdgkb/tsr-tegister-server-v1/handlers/disabilities"
	"mdgkb/tsr-tegister-server-v1/handlers/documenttypes"
	"mdgkb/tsr-tegister-server-v1/handlers/drug"
	"mdgkb/tsr-tegister-server-v1/handlers/edvs"
	"mdgkb/tsr-tegister-server-v1/handlers/fileinfos"
	"mdgkb/tsr-tegister-server-v1/handlers/hmfsescalequestions"
	"mdgkb/tsr-tegister-server-v1/handlers/humans"
	"mdgkb/tsr-tegister-server-v1/handlers/insurancecompany"
	"mdgkb/tsr-tegister-server-v1/handlers/meta"
	"mdgkb/tsr-tegister-server-v1/handlers/mkbitems"
	"mdgkb/tsr-tegister-server-v1/handlers/patientdiagnosis"
	"mdgkb/tsr-tegister-server-v1/handlers/patienthistories"
	"mdgkb/tsr-tegister-server-v1/handlers/patients"
	"mdgkb/tsr-tegister-server-v1/handlers/patientsregisters"
	"mdgkb/tsr-tegister-server-v1/handlers/patientsrepresentatives"
	"mdgkb/tsr-tegister-server-v1/handlers/patientsresearches"
	"mdgkb/tsr-tegister-server-v1/handlers/patientsresearchespools"
	"mdgkb/tsr-tegister-server-v1/handlers/questions"
	"mdgkb/tsr-tegister-server-v1/handlers/regions"
	"mdgkb/tsr-tegister-server-v1/handlers/registerpropertytouser"
	"mdgkb/tsr-tegister-server-v1/handlers/registerquery"
	"mdgkb/tsr-tegister-server-v1/handlers/registers"
	"mdgkb/tsr-tegister-server-v1/handlers/representative"
	"mdgkb/tsr-tegister-server-v1/handlers/representativetypes"
	"mdgkb/tsr-tegister-server-v1/handlers/researches"
	"mdgkb/tsr-tegister-server-v1/handlers/researchespools"
	"mdgkb/tsr-tegister-server-v1/handlers/researchesresults"
	"mdgkb/tsr-tegister-server-v1/handlers/search"
	"mdgkb/tsr-tegister-server-v1/handlers/users"
	anamnesesRouter "mdgkb/tsr-tegister-server-v1/routing/anamneses"
	authRouter "mdgkb/tsr-tegister-server-v1/routing/auth"
	chopScaleQuestionsRouter "mdgkb/tsr-tegister-server-v1/routing/chopscalequestions"
	disabilitiesRouter "mdgkb/tsr-tegister-server-v1/routing/disabilities"
	documentTypesRouter "mdgkb/tsr-tegister-server-v1/routing/documenttypes"
	drugRouter "mdgkb/tsr-tegister-server-v1/routing/drug"
	edvsRouter "mdgkb/tsr-tegister-server-v1/routing/edvs"
	fileInfoRouter "mdgkb/tsr-tegister-server-v1/routing/fileinfo"
	hmfseScaleQuestionsRouter "mdgkb/tsr-tegister-server-v1/routing/hmfsescalequestions"
	humansRouter "mdgkb/tsr-tegister-server-v1/routing/humans"
	insuranceCompanyRouter "mdgkb/tsr-tegister-server-v1/routing/insurancecompany"
	metaRouter "mdgkb/tsr-tegister-server-v1/routing/meta"
	mkbItemsRouter "mdgkb/tsr-tegister-server-v1/routing/mkbitems"
	patientDiagnosisRouter "mdgkb/tsr-tegister-server-v1/routing/patientdiagnosis"
	patientHistoriesRouter "mdgkb/tsr-tegister-server-v1/routing/patienthistories"
	patientsRouter "mdgkb/tsr-tegister-server-v1/routing/patients"
	patientsRegistersRouter "mdgkb/tsr-tegister-server-v1/routing/patientsregisters"
	patientsRepresentativesRouter "mdgkb/tsr-tegister-server-v1/routing/patientsrepresentatives"
	patientsResearchesRouter "mdgkb/tsr-tegister-server-v1/routing/patientsresearches"
	patientsResearchesPoolsRouter "mdgkb/tsr-tegister-server-v1/routing/patientsresearchespools"
	regionsRouter "mdgkb/tsr-tegister-server-v1/routing/regions"
	registerGroupRouter "mdgkb/tsr-tegister-server-v1/routing/registergroup"
	registerPropertyRouter "mdgkb/tsr-tegister-server-v1/routing/registerproperty"
	registerPropertyToUserRouter "mdgkb/tsr-tegister-server-v1/routing/registerpropertytouser"
	registerQueryRouter "mdgkb/tsr-tegister-server-v1/routing/registerquery"
	registersRouter "mdgkb/tsr-tegister-server-v1/routing/registers"
	representativeRouter "mdgkb/tsr-tegister-server-v1/routing/representative"
	representativeTypesRouter "mdgkb/tsr-tegister-server-v1/routing/representativetypes"
	researchesRouter "mdgkb/tsr-tegister-server-v1/routing/researches"
	researchesPoolsRouter "mdgkb/tsr-tegister-server-v1/routing/researchespools"
	researchesResultsRouter "mdgkb/tsr-tegister-server-v1/routing/researchesresults"
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
	registersRouter.Init(api.Group("/registers"), registers.CreateHandler(helper))
	researchesRouter.Init(api.Group("/researches"), researches.CreateHandler(helper))
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
	patientsResearchesPoolsRouter.Init(api.Group("/patients-researches-pools"), patientsresearchespools.CreateHandler(helper))
	researchesResultsRouter.Init(api.Group("/researches-results"), researchesresults.CreateHandler(helper))
	patientsResearchesRouter.Init(api.Group("/patients-researches"), patientsresearches.CreateHandler(helper))
	patientsRegistersRouter.Init(api.Group("/patients-registers"), patientsregisters.CreateHandler(helper))
	patientDiagnosisRouter.Init(api.Group("/patient-diagnosis"), patientdiagnosis.CreateHandler(helper))
	disabilitiesRouter.Init(api.Group("/disabilities"), disabilities.CreateHandler(helper))
	edvsRouter.Init(api.Group("/edvs"), edvs.CreateHandler(helper))
	humansRouter.Init(api.Group("/humans"), humans.CreateHandler(helper))
	patientsRepresentativesRouter.Init(api.Group("/patients-representatives"), patientsrepresentatives.CreateHandler(helper))
	anamnesesRouter.Init(api.Group("/anamneses"), anamneses.CreateHandler(helper))
	patientHistoriesRouter.Init(api.Group("/patient-histories"), patienthistories.CreateHandler(helper))
}
