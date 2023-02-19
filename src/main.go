package main

import (
	"example_onion/application/handler"
	"example_onion/application/service"
	"example_onion/domain/repository"
	"example_onion/infrastructure"
	"example_onion/infrastructure/database/mysql"
	"example_onion/infrastructure/persistence"
	"example_onion/infrastructure/webapi"
	"fmt"
	"net/http"
)

func init() {
	infrastructure.Load()
}

func main() {
	client := mysql.NewClient()
	defer client.Close()

	db := client.GetDb()

	userPersistence := persistence.NewUserPersistence(db)
	userRepository := repository.NewUserRepository(userPersistence)
	userService := service.NewUserService(userRepository)

	router := webapi.NewRouter(handler.HandlerPackage{UserHandler: handler.NewUserHandler(userService)}, userService)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
