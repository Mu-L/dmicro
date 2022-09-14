package main

import (
	"github.com/osgochina/dmicro/dserver"
	"github.com/osgochina/dmicro/examples/dserver/sandbox"
	"github.com/osgochina/dmicro/logger"
)

func main() {
	dserver.CloseCtl()
	dserver.SetName("DMicro")
	dserver.Setup(func(svr *dserver.DServer) {
		err := svr.AddSandBox(new(sandbox.AdminSandBox), svr.NewService("admin"))
		if err != nil {
			logger.Fatal(err)
		}
	})
}
