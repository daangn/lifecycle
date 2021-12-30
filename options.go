package lifecycle

import (
	"net"
	"net/http"

	"google.golang.org/grpc"
)

var defaultRunOptions = &runOptions{
	apps: []App{},
}

type RunOption func(o *runOptions)

type runOptions struct {
	apps []App
}

func evaluateRunOptions(opts []RunOption) *runOptions {
	optCopy := &runOptions{}
	*optCopy = *defaultRunOptions
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}

func WithGRPC(srv *grpc.Server, lis net.Listener) RunOption {
	return func(o *runOptions) {
		o.apps = append(o.apps, &grpcApp{
			srv: srv,
			lis: lis,
		})
	}
}

func WithHTTP(srv *http.Server) RunOption {
	return func(o *runOptions) {
		o.apps = append(o.apps, &httpApp{
			srv: srv,
		})
	}
}

func WithApp(app App) RunOption {
	return func(o *runOptions) {
		o.apps = append(o.apps, app)
	}
}
