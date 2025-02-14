package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"log/slog"

	"github.com/HMasataka/gofiles"
	"github.com/HMasataka/gogen"
)

const (
	ENUM_FILE_NAME     = "enums.json"
	TEMPLATE_FILE_NAME = "templates/example.tmpl"
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
	logger := gogen.GetTextLogger()

	enums, err := readEnums(ENUM_FILE_NAME)
	if err != nil {
		logger.Error("read enum files", slog.Any("error", err))
		os.Exit(1)
	}

	fmt.Println(enums)

	for _, e := range enums {
		if err := gofiles.CreateDirectoryIfNotExist(e.Package); err != nil {
			logger.Error("create directory", slog.Any("error", err))
			os.Exit(1)
		}

		path := path.Join(e.Package, fmt.Sprintf("%s.go", strings.ToLower(e.Type)))

		f, err := gofiles.CreateWriteFile(path)
		if err != nil {
			logger.Error("create file", slog.Any("error", err))
			os.Exit(1)
		}
		defer f.Close()

		enumTemplate, err := gogen.ReadTemplate(TEMPLATE_FILE_NAME)
		if err != nil {
			logger.Error("read template", slog.Any("error", err))
			os.Exit(1)
		}

		if err := enumTemplate.Execute(f, e); err != nil {
			logger.Error("execute template", slog.Any("error", err))
			os.Exit(1)
		}
	}
}
