package generator

import (
	_ "embed"
	"path/filepath"

	"github.com/wahyusa/goartisan/internal/config"
	"github.com/wahyusa/goartisan/internal/utils"
)

//go:embed templates/service.go.tmpl
var serviceTemplate []byte

func GenerateService(name string) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	servicePath := filepath.Join(cfg.App.Folder, "service", utils.ToSnakeCase(name)+"_service.go")

	data := map[string]string{
		"ModelName":   utils.ToCamelCase(name),
		"ProjectName": cfg.Module.Name,
	}

	if err := utils.ExecuteTemplateFromBytes(serviceTemplate, data, servicePath); err != nil {
		return err
	}

	return nil
}
