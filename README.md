## A Real Time Logger Client Using gRPC-Web

Following the GA announcement of gRPC-web last october, I'm creating this small project for learning purposes.

The project has the client side written in Javascript and it is statically served. The backend is Golang using NYTimes Gizmo framework.

I'm going to use Envoy to translate the HTTP/1.1 calls produced by the client into HTTP/2 calls to the server using gRPC interface.

I'm using Kubernetes for deployment. I'm also experimenting Istio and if I get good results, it will also be included here.
