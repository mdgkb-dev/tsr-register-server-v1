package routing

import (
	"mdgkb/tsr-tegister-server-v1/handlers/anamneses"
	"mdgkb/tsr-tegister-server-v1/handlers/auth"
	"mdgkb/tsr-tegister-server-v1/handlers/commissions"
	"mdgkb/tsr-tegister-server-v1/handlers/commissionsdoctors"
	"mdgkb/tsr-tegister-server-v1/handlers/commissionsdrugapplications"
	"mdgkb/tsr-tegister-server-v1/handlers/commissionstemplates"
	"mdgkb/tsr-tegister-server-v1/handlers/dataexport"
	"mdgkb/tsr-tegister-server-v1/handlers/disabilities"
	"mdgkb/tsr-tegister-server-v1/handlers/doctors"
	"mdgkb/tsr-tegister-server-v1/handlers/documentfieldvalues"
	"mdgkb/tsr-tegister-server-v1/handlers/documentfileinfos"
	"mdgkb/tsr-tegister-server-v1/handlers/documents"
	"mdgkb/tsr-tegister-server-v1/handlers/documenttypes"
	"mdgkb/tsr-tegister-server-v1/handlers/drugapplications"
	"mdgkb/tsr-tegister-server-v1/handlers/drugarrives"
	"mdgkb/tsr-tegister-server-v1/handlers/drugdecreases"
	"mdgkb/tsr-tegister-server-v1/handlers/drugdozes"
	"mdgkb/tsr-tegister-server-v1/handlers/drugforms"
	"mdgkb/tsr-tegister-server-v1/handlers/drugrecipes"
	"mdgkb/tsr-tegister-server-v1/handlers/drugs"
	"mdgkb/tsr-tegister-server-v1/handlers/edvs"
	"mdgkb/tsr-tegister-server-v1/handlers/fileinfos"
	"mdgkb/tsr-tegister-server-v1/handlers/fundcontracts"
	"mdgkb/tsr-tegister-server-v1/handlers/fundcouncils"
	"mdgkb/tsr-tegister-server-v1/handlers/humans"
	"mdgkb/tsr-tegister-server-v1/handlers/insurancecompany"
	"mdgkb/tsr-tegister-server-v1/handlers/menus"
	"mdgkb/tsr-tegister-server-v1/handlers/meta"
	"mdgkb/tsr-tegister-server-v1/handlers/mkbitems"
	"mdgkb/tsr-tegister-server-v1/handlers/patientdiagnosis"
	"mdgkb/tsr-tegister-server-v1/handlers/patienthistories"
	"mdgkb/tsr-tegister-server-v1/handlers/patients"
	"mdgkb/tsr-tegister-server-v1/handlers/patientsdomains"
	"mdgkb/tsr-tegister-server-v1/handlers/patientsrepresentatives"
	"mdgkb/tsr-tegister-server-v1/handlers/patientsresearches"
	"mdgkb/tsr-tegister-server-v1/handlers/patientsresearchespools"
	"mdgkb/tsr-tegister-server-v1/handlers/questions"
	"mdgkb/tsr-tegister-server-v1/handlers/regions"
	"mdgkb/tsr-tegister-server-v1/handlers/representatives"
	"mdgkb/tsr-tegister-server-v1/middleware"
	menusRouter "mdgkb/tsr-tegister-server-v1/routing/menus"
	representativesRouter "mdgkb/tsr-tegister-server-v1/routing/representatives"

	"mdgkb/tsr-tegister-server-v1/handlers/representativetypes"
	"mdgkb/tsr-tegister-server-v1/handlers/researches"
	"mdgkb/tsr-tegister-server-v1/handlers/researchespools"
	"mdgkb/tsr-tegister-server-v1/handlers/researchesresults"
	"mdgkb/tsr-tegister-server-v1/handlers/search"
	"mdgkb/tsr-tegister-server-v1/handlers/statuses"
	"mdgkb/tsr-tegister-server-v1/handlers/users"
	anamnesesRouter "mdgkb/tsr-tegister-server-v1/routing/anamneses"
	authRouter "mdgkb/tsr-tegister-server-v1/routing/auth"
	commissionsRouter "mdgkb/tsr-tegister-server-v1/routing/commissions"
	commissionsDoctorsRouter "mdgkb/tsr-tegister-server-v1/routing/commissionsdoctors"
	commissionsdrugapplicationsRouter "mdgkb/tsr-tegister-server-v1/routing/commissionsdrugapplications"
	commissionsTemplatesRouter "mdgkb/tsr-tegister-server-v1/routing/commissionstemplates"
	dataexportRouter "mdgkb/tsr-tegister-server-v1/routing/dataexport"
	disabilitiesRouter "mdgkb/tsr-tegister-server-v1/routing/disabilities"
	doctorsRouter "mdgkb/tsr-tegister-server-v1/routing/doctors"
	documentfieldvaluesRouter "mdgkb/tsr-tegister-server-v1/routing/documentfieldvalues"
	documentfileinfosRouter "mdgkb/tsr-tegister-server-v1/routing/documentfileinfos"
	documentsRouter "mdgkb/tsr-tegister-server-v1/routing/documents"
	documentTypesRouter "mdgkb/tsr-tegister-server-v1/routing/documenttypes"
	drugapplicationsRouter "mdgkb/tsr-tegister-server-v1/routing/drugapplications"
	drugarrivesRouter "mdgkb/tsr-tegister-server-v1/routing/drugarrives"
	drugdecreasesRouter "mdgkb/tsr-tegister-server-v1/routing/drugdecreases"
	drugdozesRouter "mdgkb/tsr-tegister-server-v1/routing/drugdozes"
	drugformsRouter "mdgkb/tsr-tegister-server-v1/routing/drugforms"
	drugrecipesRouter "mdgkb/tsr-tegister-server-v1/routing/drugrecipes"
	drugsRouter "mdgkb/tsr-tegister-server-v1/routing/drugs"
	edvsRouter "mdgkb/tsr-tegister-server-v1/routing/edvs"
	fileInfosRouter "mdgkb/tsr-tegister-server-v1/routing/fileinfos"
	fundcontractsRouter "mdgkb/tsr-tegister-server-v1/routing/fundcontracts"
	fundcouncilsRouter "mdgkb/tsr-tegister-server-v1/routing/fundcouncils"
	humansRouter "mdgkb/tsr-tegister-server-v1/routing/humans"
	insuranceCompanyRouter "mdgkb/tsr-tegister-server-v1/routing/insurancecompany"
	metaRouter "mdgkb/tsr-tegister-server-v1/routing/meta"
	mkbItemsRouter "mdgkb/tsr-tegister-server-v1/routing/mkbitems"
	patientDiagnosisRouter "mdgkb/tsr-tegister-server-v1/routing/patientdiagnosis"
	patientHistoriesRouter "mdgkb/tsr-tegister-server-v1/routing/patienthistories"
	patientsRouter "mdgkb/tsr-tegister-server-v1/routing/patients"
	patientsdomainsRouter "mdgkb/tsr-tegister-server-v1/routing/patientsdomains"
	patientsRepresentativesRouter "mdgkb/tsr-tegister-server-v1/routing/patientsrepresentatives"
	patientsResearchesRouter "mdgkb/tsr-tegister-server-v1/routing/patientsresearches"
	patientsResearchesPoolsRouter "mdgkb/tsr-tegister-server-v1/routing/patientsresearchespools"
	questionsRouter "mdgkb/tsr-tegister-server-v1/routing/questions"
	registerPropertyRouter "mdgkb/tsr-tegister-server-v1/routing/questions"
	regionsRouter "mdgkb/tsr-tegister-server-v1/routing/regions"
	representativeTypesRouter "mdgkb/tsr-tegister-server-v1/routing/representativetypes"
	researchesRouter "mdgkb/tsr-tegister-server-v1/routing/researches"
	researchesPoolsRouter "mdgkb/tsr-tegister-server-v1/routing/researchespools"
	researchesResultsRouter "mdgkb/tsr-tegister-server-v1/routing/researchesresults"
	registerGroupRouter "mdgkb/tsr-tegister-server-v1/routing/researchsection"
	searchRouter "mdgkb/tsr-tegister-server-v1/routing/search"
	drugapplicationsstatusesRouter "mdgkb/tsr-tegister-server-v1/routing/statuses"
	statusesRouter "mdgkb/tsr-tegister-server-v1/routing/statuses"
	usersRouter "mdgkb/tsr-tegister-server-v1/routing/users"

	//"mdgkb/tsr-tegister-server-v1/handlers/representative"
	"github.com/gin-gonic/gin"

	helperPack "github.com/pro-assistance/pro-assister/helper"
)

func Init(r *gin.Engine, helper *helperPack.Helper) {
	m := middleware.CreateMiddleware(helper)

	r.Static("/api/v1/static", "./static/")
	//r.Use(helper.HTTP.CORSMiddleware())
	authGroup := r.Group("/api/v1")
	api := r.Group("/api/v1")
	api.Use(m.InjectRequestInfo())
	authRouter.Init(authGroup.Group("/auth"), auth.CreateHandler(helper))
	documentTypesRouter.Init(api.Group("/document-types"), documenttypes.CreateHandler(helper))
	drugsRouter.Init(api.Group("/drugs"), drugs.CreateHandler(helper))
	fileInfosRouter.Init(api.Group("/file-infos"), fileinfos.CreateHandler(helper))
	insuranceCompanyRouter.Init(api.Group("/insurance-companies"), insurancecompany.CreateHandler(helper))
	metaRouter.Init(api.Group("/meta"), meta.CreateHandler(helper))
	mkbItemsRouter.Init(api.Group("/mkb-items"), mkbitems.CreateHandler(helper))

	registerGroupRouter.Init(api.Group("/register-groups"), questions.CreateHandler(helper))
	registerPropertyRouter.Init(api.Group("/register-properties"), questions.CreateHandler(helper))
	representativesRouter.Init(api.Group("/representatives"), representatives.CreateHandler(helper))
	representativeTypesRouter.Init(api.Group("/representative-types"), representativetypes.CreateHandler(helper))
	usersRouter.Init(api.Group("/users"), users.CreateHandler(helper))
	regionsRouter.Init(api.Group("/regions"), regions.CreateHandler(helper))
	researchesPoolsRouter.Init(api.Group("/researches-pools"), researchespools.CreateHandler(helper))
	searchRouter.Init(api.Group("/search"), search.CreateHandler(helper))
	patientsResearchesPoolsRouter.Init(api.Group("/patients-researches-pools"), patientsresearchespools.CreateHandler(helper))
	researchesResultsRouter.Init(api.Group("/researches-results"), researchesresults.CreateHandler(helper))
	patientsResearchesRouter.Init(api.Group("/patients-researches"), patientsresearches.CreateHandler(helper))
	patientDiagnosisRouter.Init(api.Group("/patient-diagnosis"), patientdiagnosis.CreateHandler(helper))
	disabilitiesRouter.Init(api.Group("/disabilities"), disabilities.CreateHandler(helper))
	edvsRouter.Init(api.Group("/edvs"), edvs.CreateHandler(helper))

	humans.Init(helper)
	humansRouter.Init(api.Group("/humans"), patientsdomains.H)

	menus.Init(helper)
	menusRouter.Init(api.Group("/menus"), menus.H)

	questions.Init(helper)
	questionsRouter.Init(api.Group("/questions"), questions.H)

	patientsdomains.Init(helper)
	patientsdomainsRouter.Init(api.Group("/patients-domains"), patientsdomains.H)

	patients.Init(helper)
	patientsRouter.Init(api.Group("/patients"), patients.H)

	researches.Init(helper)
	researchesRouter.Init(api.Group("/researches"), researches.H)

	dataexport.Init(helper)
	dataexportRouter.Init(api.Group("/data-export"), dataexport.H)

	patientsRepresentativesRouter.Init(api.Group("/patients-representatives"), patientsrepresentatives.CreateHandler(helper))
	anamnesesRouter.Init(api.Group("/anamneses"), anamneses.CreateHandler(helper))
	patientHistoriesRouter.Init(api.Group("/patient-histories"), patienthistories.CreateHandler(helper))
	commissionsTemplatesRouter.Init(api.Group("/commissions-templates"), commissionstemplates.CreateHandler(helper))
	commissionsRouter.Init(api.Group("/commissions"), commissions.CreateHandler(helper))
	commissionsDoctorsRouter.Init(api.Group("/commissions-doctors"), commissionsdoctors.CreateHandler(helper))
	doctorsRouter.Init(api.Group("/doctors"), doctors.CreateHandler(helper))
	drugarrivesRouter.Init(api.Group("/drug-arrives"), drugarrives.CreateHandler(helper))
	fundcontractsRouter.Init(api.Group("/fund-contracts"), fundcontracts.CreateHandler(helper))
	fundcouncilsRouter.Init(api.Group("/fund-councils"), fundcouncils.CreateHandler(helper))
	statusesRouter.Init(api.Group("/statuses"), statuses.CreateHandler(helper))
	drugapplicationsRouter.Init(api.Group("/drug-applications"), drugapplications.CreateHandler(helper))
	commissionsdrugapplicationsRouter.Init(api.Group("/commissions-drug-applications"), commissionsdrugapplications.CreateHandler(helper))
	drugapplicationsstatusesRouter.Init(api.Group("/drug-applications-statuses"), statuses.CreateHandler(helper))
	drugdecreasesRouter.Init(api.Group("/drug-decreases"), drugdecreases.CreateHandler(helper))
	drugrecipesRouter.Init(api.Group("/drug-recipes"), drugrecipes.CreateHandler(helper))

	documentsRouter.Init(api.Group("/documents"), documents.CreateHandler(helper))
	documentfieldvaluesRouter.Init(api.Group("/document-field-values"), documentfieldvalues.CreateHandler(helper))
	documentfileinfosRouter.Init(api.Group("/document-file-infos"), documentfileinfos.CreateHandler(helper))
	drugformsRouter.Init(api.Group("/drug-forms"), drugforms.CreateHandler(helper))
	drugdozesRouter.Init(api.Group("/drug-dozes"), drugdozes.CreateHandler(helper))
}
