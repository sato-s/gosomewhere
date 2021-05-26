package main

type Element struct {
	Name string
	URL  string
}

type HTMLList []Element

func CreateHTMLList(destinations Destinations) HTMLList {
	var list HTMLList
	for k, v := range destinations {
		list = append(list, Element{Name: k, URL: v})
	}
	return list
}
