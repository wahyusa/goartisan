package utils

import (
	"bytes"
	"os"
	"text/template"
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
