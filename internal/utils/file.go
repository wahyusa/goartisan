package utils

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/wahyusa/goartisan/internal/config"
)

func ExecuteTemplateFromBytes(templateContent []byte, data interface{}, outputPath string) error {
	// I can have custom function inside the template, toLowerCase for example is from utils/string.go
	funcMap := template.FuncMap{
		"toLowerCase": ToLowerCase,
	}

	tmpl, err := template.New("template").Funcs(funcMap).Parse(string(templateContent))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

// GenerateFile generates a file from a template with the given name and type.
// used in make_generator.go
func GenerateFile(name, fileType string, templateContent []byte) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	filePath := filepath.Join(cfg.App.Folder, fileType, ToSnakeCase(name)+"_"+fileType+".go")

	data := map[string]string{
		"ModelName":   ToCamelCase(name),
		"ProjectName": cfg.Module.Name,
	}

	return ExecuteTemplateFromBytes(templateContent, data, filePath)
}

// CreateDirectories creates multiple directories.
func CreateDirectories(basePath string, dirs []string) error {
	for _, dir := range dirs {
		dirPath := filepath.Join(basePath, dir)
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// CreateFiles creates multiple files.
func CreateFiles(basePath string, files []string) error {
	for _, filePath := range files {
		fullPath := filepath.Join(basePath, filePath)
		file, err := os.Create(fullPath)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}
