package node

import (
	channel2 "cloudCli/channel"
	"cloudCli/ctx"
	"cloudCli/gin/controller/test"
	"cloudCli/gin/routers"
	"cloudCli/utils/log"
	"context"
	"net/http"
)

type Gin struct {
	AbstractNode
}

var cliCtx ctx.Context = ctx.CreateContext(nil)

func (*Gin) Init() {
}

func (*Gin) Start(context ctx.Context) {
	routers.Include(test.Routers)
	r := routers.Init()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	cliCtx.AddAttr("gin.srv", srv)
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("listen: %s", err)
		}
	}()
}

/**
处理消息
*/
func (b *Gin) HandleMessage(msg interface{}) *AsyncResponse {
	switch msg.(type) {
	case channel2.CommandMessage:
		{
			switch msg.(*channel2.CommandMessage).Name {
			case channel2.MESSAGE_ONTIME:
				{
					//执行定时业务逻辑
				}
			}
		}
	}
	return nil
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
