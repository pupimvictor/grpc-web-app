package main

import (
	"fmt"
	"github.com/NYTimes/gizmo/config"
	"github.com/NYTimes/gizmo/server/kit"


)

func main(){
	fmt.Println("hello grpc")

	//var cfg server.Config
	config.LoadEnvConfig(nil)

	// runs the HTTP _AND_ gRPC servers
	err := kit.Run(nil)
	if err != nil {
		panic("problems running service: " + err.Error())
	}
}


