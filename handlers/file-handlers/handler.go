package fileHandlers

import (
	"github.com/imsujan276/go-clean-repo/controllers/file-controllers"
)

type handler struct {
	service filecontrollers.Service
}

func NewCreateHandler(service filecontrollers.Service) *handler {
	return &handler{service: service}
}
