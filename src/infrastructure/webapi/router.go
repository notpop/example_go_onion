package webapi

import (
	"example_onion/application/handler"
	"example_onion/application/service"
	"fmt"

	"github.com/gorilla/mux"
	"reflect"
)

func NewRouter(handlerPackage handler.HandlerPackage, userService *service.UserService) *mux.Router {
	r := mux.NewRouter()

	handlerType := reflect.TypeOf(handlerPackage)
	for i := 0; i < handlerType.NumField(); i++ {
		handler := reflect.ValueOf(handlerPackage).Field(i).Interface().(handler.Handler)
		fmt.Printf("registering %T\n", handler)
		handler.Register(r, userService)
	}

	return r
}
