package main

import (
	"myTest/handler"
	pb "myTest/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("mytest"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterMyTestHandler(srv.Server(), new(handler.MyTest))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
