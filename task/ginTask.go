package task

import (
	"cloudCli/ctx"
	"cloudCli/gin/controller/test"
	"cloudCli/gin/routers"
	"cloudCli/utils/log"
	"context"
	"net/http"
)

type Gin struct {
	AbstractTask
}

var cliCtx ctx.Context = ctx.CreateContext()

func (*Gin) Init() {
}

func (*Gin) Start(params TaskParams) {
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

func (*Gin) Stop() {

	/**
	 * 需要改为原生的http启停方式
	 * 实现优雅的关闭
	 */
	log.Info("Stop Task gin UI")
	srv := cliCtx.GetAttr("gin.srv")
	srv.(*http.Server).Shutdown(context.Background())

}
