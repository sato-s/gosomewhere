package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"io"
)

//go:embed toppage/index.html
var indexHTML string

//go:embed toppage/jquery-3.6.0.min.js
var jqueryJS template.JS

//go:embed toppage/list.min.js
var listJS template.JS

type TopPage struct {
	template     *template.Template
	destinations Destinations
}

func NewTopPage(destinations Destinations) (*TopPage, error) {
	template, err := template.New("toppage").Parse(indexHTML)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse template %+v", err)
	}
	return &TopPage{template: template, destinations: destinations}, nil
}

type Element struct {
	Name string
	URL  string
}

func (t *TopPage) list() []Element {
	var list []Element
	for k, v := range t.destinations {
		list = append(list, Element{Name: k, URL: v})
	}
	return list
}

func (t *TopPage) Execute(wr io.Writer) error {
	data := struct {
		JqueryJS template.JS
		ListJS   template.JS
		List     []Element
	}{JqueryJS: jqueryJS, ListJS: listJS, List: t.list()}
	return t.template.Execute(wr, data)
}
