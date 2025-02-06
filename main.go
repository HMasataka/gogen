package gogen

import (
	_ "embed"
	"encoding/json"
	"html/template"

	"github.com/HMasataka/gofiles"
)

type Enum struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Enums struct {
	Type    string `json:"type"`
	Package string `json:"package"`
	Data    []Enum `json:"data"`
}

func ReadTemplate(path string) (*template.Template, error) {
	b, err := gofiles.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return template.Must(template.New("").Parse(string(b))), nil
}

func ReadEnums(path string) ([]Enums, error) {
	var enums []Enums

	enumBytes, err := gofiles.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(enumBytes, &enums); err != nil {
		return nil, err
	}

	return enums, nil
}
