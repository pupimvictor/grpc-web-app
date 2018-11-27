package main

import (
	"fmt"
	"github.com/NYTimes/gizmo/config"
	"github.com/NYTimes/gizmo/server/kit"
	"github.com/pupimvictor/grpc-web-app/server"

)

func main(){
	fmt.Println("hello grpc")

	var cfg server.Config
	config.LoadEnvConfig(&cfg)

	// runs the HTTP _AND_ gRPC servers
	err := kit.Run(server.NewRPCService(&cfg))
	if err != nil {
		panic("problems running service: " + err.Error())
	}
}


