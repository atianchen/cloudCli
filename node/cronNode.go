package node

import (
	"cloudCli/cfg"
	"cloudCli/channel"
	"cloudCli/node/extend"
	"cloudCli/utils/log"
	"github.com/robfig/cron/v3"
	"reflect"
)

/**
定时任务调度
*/
type CronNode struct {
	AbstractNode
	cronInstance *cron.Cron
}

func (t *CronNode) Init() error {
	return nil
}

func (t *CronNode) Start(context *NodeContext) {
	t.cronInstance = cron.New()
	cronExpress, _ := cfg.GetConfig("cli.timer")
	if cronExpress != nil {
		for node, express := range cronExpress.(map[string]interface{}) {
			go func(node string, express string) {
				t.cronInstance.AddFunc(express, func() {
					onTimeExecute(node)
				})
			}(node, express.(string))

		}
		t.cronInstance.Start()
	}
}

func onTimeExecute(node string) {
	defer func() {
		if r := recover(); r != nil {
			log.Error("Time Execute Error:" + node)
		}
	}()
	log.Info("OnTime ", node)
	nc, _ := channel.GetChan(node)
	if nc != nil {
		nc <- channel.BulidOnTimeMessage(nil)
	}
}

func (t *CronNode) HandleMessage(msg interface{}, channel chan interface{}) {

}

func (t *CronNode) GetMsgHandler() extend.MsgHandler {
	return t
}

func (t *CronNode) Stop() {
	t.cronInstance.Stop()
}

func (t *CronNode) Name() string {
	return reflect.TypeOf(t).Elem().Name()
}
