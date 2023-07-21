package todohandlers

import todocontrollers "github.com/imsujan276/go-clean-repo/controllers/todo-controllers"

type handler struct {
	service todocontrollers.Service
}

func NewCreateHandler(service todocontrollers.Service) *handler {
	return &handler{service: service}
}
