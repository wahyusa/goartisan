package generator

import (
	_ "embed"
	"path/filepath"

	"github.com/wahyusa/goartisan/internal/config"
	"github.com/wahyusa/goartisan/internal/utils"
)

//go:embed templates/model.go.tmpl
var modelTemplate []byte

func GenerateModel(name string) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	modelPath := filepath.Join(cfg.App.Folder, "model", utils.ToSnakeCase(name)+".go")

	data := map[string]string{
		"ModelName": utils.ToCamelCase(name),
	}

	if err := utils.ExecuteTemplateFromBytes(modelTemplate, data, modelPath); err != nil {
		return err
	}

	return nil
}
