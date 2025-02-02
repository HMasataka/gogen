package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path"
	"strings"

	"github.com/HMasataka/gofiles"
)

const (
	ENUM_FILE_NAME     = "enums.json"
	TEMPLATE_FILE_NAME = "main.tmpl"
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

func readTemplate(path string) *template.Template {
	b, err := gofiles.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return template.Must(template.New("").Parse(string(b)))
}

func readEnums(path string) ([]Enums, error) {
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

func main() {
	enums, err := readEnums(ENUM_FILE_NAME)
	if err != nil {
		panic(err)
	}

	fmt.Println(enums)

	for _, e := range enums {
		if err := os.Mkdir(e.Package, 0755); err != nil {
			panic(err)
		}

		path := path.Join(e.Package, fmt.Sprintf("%s.go", strings.ToLower(e.Type)))

		f, err := gofiles.CreateWriteFile(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		enumTemplate := readTemplate(TEMPLATE_FILE_NAME)

		if err := enumTemplate.Execute(f, e); err != nil {
			panic(err)
		}
	}
}
