package gin

import (
	"cloudCli/cfg"
	"cloudCli/ctx"
	"cloudCli/gin/controller"
	"cloudCli/gin/controller/sysAction"
	"cloudCli/gin/controller/test"
	"cloudCli/gin/routers"
	"cloudCli/node"
	"cloudCli/utils/log"
	"context"
	"net/http"
)

type Gin struct {
	node.AbstractNode
}

var cliCtx ctx.Context = node.CreateNodeContext(nil)
var actions = []controller.WebAction{sysAction.SysAction{}}

func (*Gin) Init() {
}

func (*Gin) Start(context *node.NodeContext) {
	port, err := cfg.GetConfig("cli.server.port")
	if err == nil {
		routers.Include(test.Routers)
		for _, action := range actions {
			action.InitAction()
			routers.Include(action.AddRouter)
		}

		r := routers.Init()
		srv := &http.Server{
			Addr:    ":" + port.(string),
			Handler: r,
		}
		cliCtx.AddAttr("gin.srv", srv)
		go func() {
			// 服务连接
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Errorf("listen: %s", err)
			}
		}()
	} else {
		log.Error("Start Gin Error ", err)
	}

}

func (d *Gin) HandleMessage(msg interface{}) {

}

func (d *Gin) GetMsgHandler() node.MsgHandler {
	return d
}

func (*Gin) Stop() {

	/**
	 * 需要改为原生的http启停方式
	 * 实现优雅的关闭
	 */
	log.Info("Stop Task gin UI")
	srv := cliCtx.GetAttr("gin.srv")
	srv.(*http.Server).Shutdown(context.Background())

}
