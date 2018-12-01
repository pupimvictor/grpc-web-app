package server

import (
	"github.com/NYTimes/gizmo/server/kit"
	"github.com/go-kit/kit/endpoint"
	"net/http"

	"cloud.google.com/go/datastore"
	"context"
	"github.com/NYTimes/gizmo/pubsub"
	"github.com/NYTimes/gziphandler"
	httptransport "github.com/go-kit/kit/transport/http"
	"google.golang.org/grpc"
)

type (
	// service will implement server.service and
	// handle all requests to the server.
	service struct {
		inputStream       pubsub.Subscriber
		ds                *datastore.Client
		passthroughCh     chan chan *Event
		stopPassthroughCh chan bool
	}
	// Config is a struct to contain all the needed
	// configuration for our service
	Config struct {

	}
)

// New will instantiate a service
// with the given configuration.
func New(cfg *Config) kit.Service{
	//todo init ds and p/s
	return &service{


	}
}

func (s service) HTTPRouterOptions() []kit.RouterOption {
	return nil
}

func (s service) HTTPOptions() []httptransport.ServerOption {
	return nil
}

// HTTPMiddleware provides an http.Handler hook wrapped around all requests.
// In this implementation, we're using a GzipHandler middleware to
// compress our responses.
func (s service) HTTPMiddleware(h http.Handler) http.Handler {
	return gziphandler.GzipHandler(h)
}

// Middleware provides an http.Handler hook wrapped around all requests.
// In this implementation, we're using a GzipHandler middleware to
// compress our responses.
func (s service) Middleware(e endpoint.Endpoint) endpoint.Endpoint {
	return e
}

// JSONEndpoints is a listing of all endpoints available in the Service.
// If using Cloud Endpoints, this is not needed but handy for local dev.
func (s service) HTTPEndpoints() map[string]map[string]kit.HTTPEndpoint {
	return map[string]map[string]kit.HTTPEndpoint{
		"/health": {
			"GET": {
				Endpoint:s.Health,
			},
		},
	}
}

func (s service) RPCMiddleware() grpc.UnaryServerInterceptor {
	return nil
}

func (s service) RPCOptions() []grpc.ServerOption {
	return nil
}

func (s service) RPCServiceDesc() *grpc.ServiceDesc {
	// snagged from the pb.go file
	return &_EventLogger_serviceDesc
}

func (s service) Health(ctx context.Context, _ interface{}) (interface{}, error) {
	return "healthy", nil
}
