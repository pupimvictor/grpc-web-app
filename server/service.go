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
	// Service will implement server.Service and
	// handle all requests to the server.
	Service struct {
		InputStream       pubsub.Subscriber
		NewStreamId       func() int64
		ds                *datastore.Client
		passthroughCh     chan chan *StreamEventsResponse
		stopPassthroughCh chan int64
	}
	// Config is a struct to contain all the needed
	// configuration for our Service
	Config struct {

	}
)

// New will instantiate a Service
// with the given configuration.
func New(cfg *Config) kit.Service{
	//todo init ds and p/s
	return &Service{


	}
}

func (s Service) HTTPRouterOptions() []kit.RouterOption {
	return nil
}

func (s Service) HTTPOptions() []httptransport.ServerOption {
	return nil
}

// HTTPMiddleware provides an http.Handler hook wrapped around all requests.
// In this implementation, we're using a GzipHandler middleware to
// compress our responses.
func (s Service) HTTPMiddleware(h http.Handler) http.Handler {
	return gziphandler.GzipHandler(h)
}

// Middleware provides an http.Handler hook wrapped around all requests.
// In this implementation, we're using a GzipHandler middleware to
// compress our responses.
func (s Service) Middleware(e endpoint.Endpoint) endpoint.Endpoint {
	return e
}

// JSONEndpoints is a listing of all endpoints available in the Service.
// If using Cloud Endpoints, this is not needed but handy for local dev.
func (s Service) HTTPEndpoints() map[string]map[string]kit.HTTPEndpoint {
	return map[string]map[string]kit.HTTPEndpoint{
		"/health": {
			"GET": {
				Endpoint:s.Health,
			},
		},
	}
}

func (s Service) RPCMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){

		return
	}
}

func (s Service) RPCOptions() []grpc.ServerOption {
	return nil
}

func (s Service) RPCServiceDesc() *grpc.ServiceDesc {
	// snagged from the pb.go file
	return &_EventLogger_serviceDesc
}

func (s Service) Health(ctx context.Context, _ interface{}) (interface{}, error) {
	return "healthy", nil
}
