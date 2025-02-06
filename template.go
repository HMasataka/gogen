package gogen

import (
	_ "embed"
	"html/template"

	"github.com/HMasataka/gofiles"
	"github.com/Masterminds/sprig/v3"
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

	return template.Must(template.New(opts.Name).Funcs(sprig.FuncMap()).Parse(string(b))), nil
}
