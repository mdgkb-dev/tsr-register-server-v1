package main

import (
	"flag"
	"fmt"
	"log"
	"mdgkb/tsr-tegister-server-v1/cmd/cli/nameBuilder"
	"mdgkb/tsr-tegister-server-v1/cmd/cli/templatesFabrics/model"
	"os"
)

func main() {
	mode := flag.String("mode", "", "init/create/createSql/run/rollback")
	action := flag.String("action", "", "init/create/createSql/run/rollback")
	name := flag.String("name", "", "init/create/createSql/run/rollback")
	flag.Parse()
	names := nameBuilder.GetNames(name)
	switch *mode {
	case "model":
		doActionModel(action, names)
	case "api":
		doActionApi(action, names)
	case "service":
		doActionService(action, names)
	}
}

func createModel(names *nameBuilder.NameFormats) {
	path := fmt.Sprintf("models/%s.go", *names.PascalCase)
	writeFile(path, model.Fabric(names))
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
