package main

import (
	"first-project/internal/configs"
	"first-project/internal/domain"
	"first-project/internal/infra/database"
	"first-project/internal/infra/database/repository"
	"first-project/internal/infra/web/server"
	"first-project/internal/usecases"
	"net/http"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := database.NewDB(config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		return
	}

	userRepository := repository.NewUserRepository(db)
	createUserUseCase := usecases.NewCreateUserUseCase(userRepository)

	webServer := server.NewWebServer(config.AppPort)
	webServer.AddHandler("/users", http.MethodPost, createUserUseCase.Execute)
	webServer.Start()
}
