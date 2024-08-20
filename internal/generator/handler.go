package generator

import (
	_ "embed"
	"path/filepath"

	"github.com/wahyusa/goartisan/internal/config"
	"github.com/wahyusa/goartisan/internal/utils"
)

//go:embed templates/handler.go.tmpl
var handlerTemplate []byte

func GenerateHandler(name string) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	handlerPath := filepath.Join(cfg.App.Folder, "handler", utils.ToSnakeCase(name)+"_handler.go")

	data := map[string]string{
		"ModelName":   utils.ToCamelCase(name),
		"ProjectName": cfg.Module.Name,
	}

	if err := utils.ExecuteTemplateFromBytes(handlerTemplate, data, handlerPath); err != nil {
		return err
	}

	return nil
}
