package main

import (
	"log"
	"os"
	"text/template"
)

type Note struct {
	Title       string
	Description string
}






const tmpl = `Note are
{{ range .}}
Title : {{ .Title }}, Description: {{ .Description }}
{{ end }}
`

func check(err error) {
	if err != nil {
		log.Fatal("Parse: ", err)
	}

}

func main() {
	notes := []Note{
        {"text/template", "Template Generate txt "},
        {"html/template", "Template Generate HTML"},
    }

	t := template.New("note")

	t, err := t.Parse(tmpl)
	check(err)
	if err := t.Execute(os.Stdout, notes); err != nil {
		log.Fatal("Excute:  ", err)
		return
	}

}
