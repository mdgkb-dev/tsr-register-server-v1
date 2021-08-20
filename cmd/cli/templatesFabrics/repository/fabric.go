package repository

import (
	"mdgkb/tsr-tegister-server-v1/cmd/cli/config"
	"mdgkb/tsr-tegister-server-v1/cmd/cli/nameBuilder"
	"mdgkb/tsr-tegister-server-v1/cmd/cli/templatesFabrics"
)

func Fabric(name *nameBuilder.NameFormats) *string {
	data := templatesFabrics.CreateData(name)
	return templatesFabrics.ParseTemplate(config.GetTemplatePath("repository"), data)
}
