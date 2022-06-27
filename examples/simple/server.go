package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/osgochina/dmicro/drpc"
	"github.com/osgochina/dmicro/drpc/plugin/ignorecase"
	"github.com/osgochina/dmicro/utils/graceful"
	"time"
)

func main() {
	go graceful.GraceSignal()
	svr := drpc.NewEndpoint(drpc.EndpointConfig{
		CountTime:   true,
		LocalIP:     "127.0.0.1",
		ListenPort:  9091,
		PrintDetail: true,
	}, ignorecase.NewIgnoreCase())

	svr.RouteCall(new(Math))

	// broadcast per 5s
	go func() {
		for {
			time.Sleep(time.Second * 5)
			svr.RangeSession(func(sess drpc.Session) bool {
				sess.Push(
					"/push/status",
					fmt.Sprintf("this is a broadcast, server time: %v", time.Now()),
				)
				return true
			})
		}
	}()

	svr.ListenAndServe()
}

type Math struct {
	drpc.CallCtx
}

func (m *Math) Add(arg *[]int) (int, *drpc.Status) {
	// test meta
	glog.Infof(context.TODO(), "author: %s", m.PeekMeta("author"))
	// add
	var r int
	for _, a := range *arg {
		r += a
	}
	// response
	return r, nil
}
