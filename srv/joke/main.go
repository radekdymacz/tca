package main

import (
	"log"

	micro "github.com/micro/go-micro"
	"github.com/radekdymacz/tca/srv/joke/handler"
	proto "github.com/radekdymacz/tca/srv/joke/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.joke"),
		micro.Version("latest"),
	)

	service.Init()
	proto.RegisterJokeHandler(service.Server(), new(handler.Joker))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
