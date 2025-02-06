package gogen

import (
	_ "embed"
	"html/template"

	"github.com/HMasataka/gofiles"
)

type ReadTemplateOptions struct {
	Name string
}

type ReadTemplateOption func(*ReadTemplateOptions)

func ReadTemplate(path string, setters ...ReadTemplateOption) (*template.Template, error) {
	opts := &ReadTemplateOptions{
		Name: "",
	}

	for _, setter := range setters {
		setter(opts)
	}

	b, err := gofiles.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return template.Must(template.New(opts.Name).Parse(string(b))), nil
}
