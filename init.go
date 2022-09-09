package main

import (
	"flag"
	"fmt"
	"master-user/config"
	"master-user/modules/share"
	"master-user/modules/v1/users/mapper"
	"master-user/modules/v1/users/presenter"
	"master-user/modules/v1/users/repository"
	"master-user/modules/v1/users/usecase"
	"os"

	"github.com/joho/godotenv"
)

type Service struct {
	UserHandler *presenter.HttpHandlerUser
}

func RunApplication() {
	isLocalDevelopment := flag.Bool("local", true, "=(true/false)")
	flag.Parse()
	if *isLocalDevelopment {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println(".env is not loaded properly")
			os.Exit(1)
		}
	}

	WritePostgres, _ := config.WritePostgres()
	ReadPostgres, _ := config.ReadPostgres()
	db := share.NewRepository(WritePostgres, ReadPostgres)

	userRepository := repository.NewUserRepositoryImpl(db)
	userMapper := mapper.NewUserMapperImpl()
	userUsecase := usecase.NewUserUsecaseImpl(userRepository, userMapper)
	userHandler := presenter.NewHttpHandlerUser(userUsecase)

	service := &Service{
		UserHandler: userHandler,
	}

	service.RoutingAndListen()
}
