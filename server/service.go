package server

import (
	"net/http"

	"github.com/NYTimes/gizmo/server"
	"google.golang.org/grpc"

	"context"
	"github.com/NYTimes/gizmo/pubsub"
	"cloud.google.com/go/datastore"
)

type (
	// service will implement server.service and
	// handle all requests to the server.
	service struct {
		inputStream        pubsub.Subscriber
		ds                 *datastore.Client
		PassthroughCh      chan chan *Event
		StopPassthroughCh  chan bool
	}
	// Config is a struct to contain all the needed
	// configuration for our service
	Config struct {
		Server *server.Config
	}
)

// NewRPCService will instantiate a service
// with the given configuration.
func NewRPCService(cfg *Config) *service {
	//todo init ds and p/s
	return &service{


	}
}

// Prefix returns the string prefix used for all endpoints within
// this service.
func (s *service) Prefix() string {
	return "/svc"
}

// Service provides the service with a description of the
// service to serve and the implementation.
func (s *service) Service() (*grpc.ServiceDesc, interface{}) {
	return &_EventLogger_serviceDesc, s
}

// Middleware provides an http.Handler hook wrapped around all requests.
// In this implementation, we're using a GzipHandler middleware to
// compress our responses.
func (s *service) Middleware(h http.Handler) http.Handler {
	return h
}

// ContextMiddleware provides a server.ContextHAndler hook wrapped around all
// requests. This could be handy if you need to decorate the request context.
func (s *service) ContextMiddleware(h server.ContextHandler) server.ContextHandler {
	return h
}

// ContextEndpoints may be needed if your server has any non-RPC-able
// endpoints. In this case, we have none but still need this method to
// satisfy the server.service interface.
func (s *service) ContextEndpoints() map[string]map[string]server.ContextHandlerFunc {
	return map[string]map[string]server.ContextHandlerFunc{}
}

// JSONContextEndpoints is a listing of all endpoints available in the service.
func (s *service) JSONEndpoints() map[string]map[string]server.JSONContextEndpoint {
	return map[string]map[string]server.JSONContextEndpoint{
		"/health": {
			"GET": s.Health,
		},
	}
}

func (s *service) Health(ctx context.Context, r *http.Request) (int, interface{}, error) {
	return http.StatusOK, "healthy", nil
}
