package main

import (
	"log"

	micro "github.com/micro/go-micro"
	"github.com/radekdymacz/tca/srv/name/handler"
	proto "github.com/radekdymacz/tca/srv/name/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.name"),
		micro.Version("latest"),
	)

	service.Init()
	proto.RegisterNameHandler(service.Server(), new(handler.Namer))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
