# Lifecycle

Golang application lifecycle manager.

**Note: This is not an official Daangn project**

# Usage

```go
package main

import (
	"net"
	"net/http"

	"github.com/daangn/lifecycle"
	"google.golang.org/grpc"
)

type yourWorkerApp interface{
	Start() error
	GracefulStop()
}

func main() {
	var (
		worker     yourWorkerApp
		httpServer *http.Server
		grpcServer *grpc.Server
		grpcLis    net.Listener
	)

	// ...

	if err := lifecycle.Run(
		lifecycle.WithGRPC(grpcServer, grpcLis),
		lifecycle.WithHTTP(httpServer),
		lifecycle.WithApp(worker),
	); err != nil {
		panic(err)
	}
}
```
