package handler

import (
	"example_onion/application/service"
	"github.com/gorilla/mux"
)

type Handler interface {
	Register(r *mux.Router, userService *service.UserService)
}

type HandlerPackage struct {
	UserHandler *UserHandler
}
