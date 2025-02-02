package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"strings"
)

const (
	ENUM_FILE_NAME = "enums.json"
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

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func readTemplate(path string) *template.Template {
	b, err := readFile(path)
	if err != nil {
		panic(err)
	}

	return template.Must(template.New("").Parse(string(b)))
}

func main() {
	var enums []Enums

	enumBytes, err := readFile(ENUM_FILE_NAME)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(enumBytes, &enums); err != nil {
		panic(err)
	}

	fmt.Println(enums)

	for _, e := range enums {
		if err := os.Mkdir(e.Package, 0755); err != nil {
			panic(err)
		}

		f, err := os.Create("./" + e.Package + "/" + strings.ToLower(e.Type) + ".go")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		enumTemplate := readTemplate("main.tmpl")

		if err := enumTemplate.Execute(f, e); err != nil {
			panic(err)
		}
	}
}
