package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"io"
)

//go:embed toppage/index.html
var indexHTML string

type TopPage struct {
	template *template.Template
}

func NewTopPage() (*TopPage, error) {
	template, err := template.New("toppage").Parse(indexHTML)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse template %+v", err)
	}
	return &TopPage{template: template}, nil
}

func (t *TopPage) Execute(wr io.Writer) error {
	return t.template.Execute(wr, struct{}{})
}
