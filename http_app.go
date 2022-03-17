// Copyright 2022 Danggeun Market Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
