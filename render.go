package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// renderTemplate will render utility
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template")
		return
	}
}
