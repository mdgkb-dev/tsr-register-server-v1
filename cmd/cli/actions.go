package main

import (
	"fmt"
	"log"
	"mdgkb/tsr-tegister-server-v1/cmd/cli/nameBuilder"
	"mdgkb/tsr-tegister-server-v1/cmd/cli/templatesFabrics/handler"
	initFabric "mdgkb/tsr-tegister-server-v1/cmd/cli/templatesFabrics/init"
	"mdgkb/tsr-tegister-server-v1/cmd/cli/templatesFabrics/model"
	"mdgkb/tsr-tegister-server-v1/cmd/cli/templatesFabrics/repository"
	"mdgkb/tsr-tegister-server-v1/cmd/cli/templatesFabrics/routing"
	"mdgkb/tsr-tegister-server-v1/cmd/cli/templatesFabrics/service"
	"os"
)

func doActionApi(action *string, names *nameBuilder.NameFormats) {
	switch *action {
	case "create":
		createModel(names)

		err := os.Mkdir(fmt.Sprintf("handlers/%s", *names.CamelCase), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		writeFile(fmt.Sprintf("handlers/%s/init.go", *names.CamelCase), initFabric.Fabric(names))
		writeFile(fmt.Sprintf("handlers/%s/handler.go", *names.CamelCase), handler.Fabric(names))
		writeFile(fmt.Sprintf("handlers/%s/service.go", *names.CamelCase), service.Fabric(names))
		writeFile(fmt.Sprintf("handlers/%s/repository.go", *names.CamelCase), repository.Fabric(names))

		err = os.Mkdir(fmt.Sprintf("routing/%s", *names.CamelCase), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		writeFile(fmt.Sprintf("routing/%s/init.go", *names.CamelCase), routing.Fabric(names))
	}
}

func doActionModel(action *string, name *nameBuilder.NameFormats) {
	switch *action {
	case "create":
		res := model.Fabric(name)
		path := fmt.Sprintf("models/%s.go", *name.PascalCase)
		writeFile(path, res)
	}
}

func doActionService(action *string, names *nameBuilder.NameFormats) {
	switch *action {
	case "create":
		createModel(names)

		err := os.Mkdir(fmt.Sprintf("handlers/%s", *names.CamelCase), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		writeFile(fmt.Sprintf("handlers/%s/init.go", *names.CamelCase), initFabric.Fabric(names))
		writeFile(fmt.Sprintf("handlers/%s/handler.go", *names.CamelCase), handler.Fabric(names))
		writeFile(fmt.Sprintf("handlers/%s/service.go", *names.CamelCase), service.Fabric(names))
		writeFile(fmt.Sprintf("handlers/%s/repository.go", *names.CamelCase), repository.Fabric(names))

		err = os.Mkdir(fmt.Sprintf("routing/%s", *names.CamelCase), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		writeFile(fmt.Sprintf("routing/%s/init.go", *names.CamelCase), routing.Fabric(names))
	}
}
