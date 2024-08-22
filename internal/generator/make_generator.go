package generator

import (
	_ "embed"

	"github.com/wahyusa/goartisan/internal/utils"
)

//go:embed templates/model.go.tmpl
var modelTemplate []byte

//go:embed templates/repository.go.tmpl
var repositoryTemplate []byte

//go:embed templates/service.go.tmpl
var serviceTemplate []byte

//go:embed templates/handler.go.tmpl
var handlerTemplate []byte

func GenerateModel(name string) error {
	return utils.GenerateFile(name, "model", modelTemplate)
}

func GenerateRepository(name string) error {
	return utils.GenerateFile(name, "repository", repositoryTemplate)
}

func GenerateService(name string) error {
	return utils.GenerateFile(name, "service", serviceTemplate)
}

func GenerateHandler(name string) error {
	return utils.GenerateFile(name, "handler", handlerTemplate)
}
