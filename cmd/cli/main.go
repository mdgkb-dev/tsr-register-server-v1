package main

import (
	"flag"
	"fmt"
	"log"
	"mdgkb/tsr-tegister-server-v1/cmd/cli/handler"
	"mdgkb/tsr-tegister-server-v1/cmd/cli/model"
	"mdgkb/tsr-tegister-server-v1/cmd/cli/repository"
	"mdgkb/tsr-tegister-server-v1/cmd/cli/routing"
	"mdgkb/tsr-tegister-server-v1/cmd/cli/strcase"
	"os"
)

func main() {
	mode := flag.String("mode", "", "init/create/createSql/run/rollback")
	action := flag.String("action", "", "init/create/createSql/run/rollback")
	name := flag.String("name", "", "init/create/createSql/run/rollback")
	flag.Parse()
	switch *mode {
	case "model":
		doActionModel(action, GetNames(name).PascalCase)
	case "api":
		doActionApi(action, GetNames(name))
	}
}

func doActionApi(action *string, name NameFormats) {
	switch *action {
	case "create":
		res := model.Fabric(&name.PascalCase, &name.CamelCase)
		path := fmt.Sprintf("models/%s.go", name.PascalCase)
		writeFile(path, &res)

		err := os.Mkdir(fmt.Sprintf("handlers/%s", name.CamelCase), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		res = handler.Fabric(&name.PascalCase, &name.CamelCase)
		writeFile(fmt.Sprintf("handlers/%s/handler.go", name.CamelCase), &res)

		res = repository.Fabric(&name.PascalCase, &name.CamelCase)
		writeFile(fmt.Sprintf("handlers/%s/repository.go", name.CamelCase), &res)

		err = os.Mkdir(fmt.Sprintf("routing/%s", name.CamelCase), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		res = routing.Fabric(&name.PascalCase, &name.CamelCase)
		writeFile(fmt.Sprintf("routing/%s/init.go", name.CamelCase), &res)
	}
}

func doActionModel(action *string, name string) {
	switch *action {
	case "create":
		res := model.Fabric(&name, &name)
		path := fmt.Sprintf("models/%s.go", name)
		writeFile(path, &res)
	}
}

func writeFile(path string, data *string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString(*data)
	if err != nil {
		log.Fatal(err)
	}
}

type NameFormats struct {
	CamelCase  string
	PascalCase string
}

func GetNames(name *string) NameFormats {
	return NameFormats{
		CamelCase:  strcase.ToLowerCamel(*name),
		PascalCase: strcase.ToCamel(*name),
	}
}
