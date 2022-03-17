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
