package main

import (
	"io"
	"text/template"
    "io/ioutil"
)

type Metadata struct {
	PackageName   string
	TypeName      string
    FullTypeName  string
}

type Generator struct {
	templateFileName string
}

// Generator factory
func NewGenerator(templateFileName string) *Generator {
    return &Generator{templateFileName:    templateFileName}
}

func (g *Generator) Generate(writer io.Writer, metadata Metadata) error {
	tmpl, err := g.template()
	if err != nil {
		return nil
	}

	return tmpl.Execute(writer, metadata)
}

func (g *Generator) template() (*template.Template, error) {

	resource, err := ioutil.ReadFile(g.templateFileName)
    if err != nil {
		return nil, err
	}

	tmpl := template.New("template")
	return tmpl.Parse(string(resource))
}
