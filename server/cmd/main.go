package main

import (
	"context"
	"fmt"
	"github.com/NYTimes/gizmo/config"
	"github.com/NYTimes/gizmo/server/kit"
	"github.com/pupimvictor/grpc-web-app/server"

)

func main(){
	fmt.Println("hello grpc")

	var cfg server.Config
	config.LoadEnvConfig(&cfg)

	s := server.New(&cfg)
	service := s.(*server.Service)
	eventsInputStream := service.InputStream.Start()

	ctx := context.Background()
	service.StartInputstream(ctx, eventsInputStream)

	// runs the HTTP _AND_ gRPC servers
	err := kit.Run(s)
	if err != nil {
		panic("problems running service: " + err.Error())
	}
}


