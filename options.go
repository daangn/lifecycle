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
