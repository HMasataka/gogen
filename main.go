package gogen

import (
	_ "embed"
	"html/template"

	"github.com/HMasataka/gofiles"
)

func ReadTemplate(path string) (*template.Template, error) {
	b, err := gofiles.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return template.Must(template.New("").Parse(string(b))), nil
}
