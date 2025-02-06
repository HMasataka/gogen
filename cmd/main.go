package main

import (
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
	TEMPLATE_FILE_NAME = "main.tmpl"
)

func main() {
	logger := gogen.GetTextLogger()

	enums, err := gogen.ReadEnums(ENUM_FILE_NAME)
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
