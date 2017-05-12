package main

import (
	"log"

	micro "github.com/micro/go-micro"
	"github.com/radekdymacz/tca/api/handler"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.joke"),
	)

	service.Init()
	service.Server().Handle(
		service.Server().NewHandler(
			&handler.Joke{service.Client()},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
