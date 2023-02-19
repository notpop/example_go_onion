package webapi

import (
	"example_onion/application/handler"
	"example_onion/application/service"
	"fmt"
	"github.com/gorilla/mux"
	"reflect"
)

type HandlerPackage struct {
	UserHandler handler.UserHandler
}

func NewRouter(handlerPackage *HandlerPackage, userService *service.UserService) *mux.Router {
	r := mux.NewRouter()

	handlerValue := reflect.ValueOf(handlerPackage)
	for i := 0; i < handlerValue.NumField(); i++ {
		fieldValue := handlerValue.Field(i)
		if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() {
			handler := fieldValue.Elem().Interface().(handler.Handler)
			fmt.Printf("registering %T\n", handler)
			handler.Register(r, userService)
		}
	}

	return r
}
