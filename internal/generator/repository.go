package generator

import (
	_ "embed"
	"path/filepath"

	"github.com/wahyusa/goartisan/internal/config"
	"github.com/wahyusa/goartisan/internal/utils"
)

//go:embed templates/repository.go.tmpl
var repositoryTemplate []byte

func GenerateRepository(name string) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	repoPath := filepath.Join(cfg.App.Folder, "repository", utils.ToSnakeCase(name)+"_repository.go")

	data := map[string]string{
		"ModelName":   utils.ToCamelCase(name),
		"ProjectName": cfg.Module.Name,
	}

	if err := utils.ExecuteTemplateFromBytes(repositoryTemplate, data, repoPath); err != nil {
		return err
	}

	return nil
}
