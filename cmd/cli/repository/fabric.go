package repository

import (
	"log"
	"strings"
	"text/template"
)

type Data struct {
	Model   string
	Package string
}

func Fabric(pascalName *string, camelName *string) string {
	var buf strings.Builder
	data := Data{Model: *pascalName, Package: *camelName}
	tmpl, err := template.ParseFiles("./cmd/cli/repository/templates/repositoryTemplate.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	_ = tmpl.Execute(&buf, data)
	return buf.String()
}
