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
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
)

type App interface {
	Start() error
	GracefulStop()
}

func Run(opts ...RunOption) (err error) {
	o := evaluateRunOptions(opts)
	c := newController()
	return c.Run(o.apps...)
}

type controller struct {
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc

	errch chan error
}

func newController() *controller {
	ctx, cancel := context.WithCancel(context.Background())

	return &controller{
		ctx:    ctx,
		cancel: cancel,
		errch:  make(chan error, 1),
	}
}

func (c *controller) Run(apps ...App) error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	for _, app := range apps {
		c.wg.Add(1)
		go c.runApp(app)
	}

	var err error
	select {
	case <-sigCh:
	case err = <-c.errch:
	}

	signal.Stop(sigCh)
	c.cancel()
	c.wg.Wait()

	return err
}

func (c *controller) runApp(app App) {
	defer c.wg.Done()

	done := uint32(0)

	go func() {
		err := app.Start()
		atomic.StoreUint32(&done, 1)
		if err != nil {
			select {
			case c.errch <- err:
			default:
			}
		}
	}()

	<-c.ctx.Done()
	if atomic.LoadUint32(&done) != 1 {
		app.GracefulStop()
	}
}
