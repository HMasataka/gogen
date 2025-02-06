package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"log/slog"

	"github.com/HMasataka/gofiles"
	"github.com/HMasataka/gogen"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Output          string `short:"o" long:"output" description:"Output file name" required:"true"`
	OutputDirectory string `short:"d" long:"output-dir" description:"Output directory path" default:"outputs"`
	Template        string `short:"t" long:"template" description:"Input template file path" required:"true"`
}

type Data struct {
	Package    string
	EntityName string
}

func main() {
	logger := gogen.GetTextLogger()

	var opts Options

	if _, err := flags.Parse(&opts); err != nil {
		logger.Error("parse args", slog.Any("error", err))
		os.Exit(1)
	}

	if err := gofiles.CreateDirectoryIfNotExist(opts.OutputDirectory); err != nil {
		logger.Error("create directory", slog.Any("error", err))
		os.Exit(1)
	}

	path := path.Join(opts.OutputDirectory, fmt.Sprintf("%s.go", strings.ToLower(opts.Output)))

	f, err := gofiles.CreateWriteFile(path)
	if err != nil {
		logger.Error("create file", slog.Any("error", err))
		os.Exit(1)
	}
	defer f.Close()

	template, err := gogen.ReadTemplate(opts.Template)
	if err != nil {
		logger.Error("read template", slog.Any("error", err))
		os.Exit(1)
	}

	data := Data{
		Package:    "driver",
		EntityName: "user",
	}

	if err := template.Execute(f, data); err != nil {
		logger.Error("execute template", slog.Any("error", err))
		os.Exit(1)
	}
}
