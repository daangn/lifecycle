package lifecycle

import (
	"context"
	"net/http"
)

var _ App = (*httpApp)(nil)

type httpApp struct {
	srv *http.Server
}

func (app *httpApp) Start() error {
	if err := app.srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (app *httpApp) GracefulStop() {
	// return only context.Err()
	_ = app.srv.Shutdown(context.Background())
}
