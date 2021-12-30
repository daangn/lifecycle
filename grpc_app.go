package lifecycle

import (
	"net"

	"google.golang.org/grpc"
)

var _ App = (*grpcApp)(nil)

type grpcApp struct {
	srv *grpc.Server
	lis net.Listener
}

func (app *grpcApp) Start() error {
	if err := app.srv.Serve(app.lis); err != nil {
		return err
	}
	return nil
}

func (app *grpcApp) GracefulStop() {
	app.srv.GracefulStop()
}
