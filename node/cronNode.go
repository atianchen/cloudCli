package node

import (
	"cloudCli/cfg"
	"cloudCli/channel"
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

func (t *CronNode) Init() {
	/**
	 * 从配置脚本加载，包括插件配置、定时的配置等
	 * 需要根据 params的内容，来决定执行那些PLUGIN
	 */
	/*	t.PluginList = []plugin.Plugin{&plugin.InspectPlugin{}, &plugin.MailPlugin{}}

		pluginParams := plugin.ExecuteParams{}
		cron := cfg.GetConfig("cli.inspect.cron")
		if cron != nil {
			log.Infof("Cron %s", cron)
			_, err := cronInstance.AddFunc("* * * * *", func() {
				for _, instance := range t.PluginList {

					instance.Execute(ctx.CreateNodeContext(instance), pluginParams)
				}
			})
			if err != nil {
				log.Info(err)
			}
		}*/
}

func (t *CronNode) Start(context *NodeContext) {
	t.cronInstance = cron.New()
	cronExpress, _ := cfg.GetConfig("cli.timer")
	if cronExpress != nil {
		for node, express := range cronExpress.(map[string]interface{}) {
			entryId, err := t.cronInstance.AddFunc(express.(string), func() {
				/**
				定时任务通知
				*/
				onTimeExecute(node)
			})
			if err != nil {
				log.Error("Timer Start Error:", err.Error(), node, express)
			} else {
				log.Info("Timer Init Success:", entryId, node, express)
			}
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
	nc := channel.GetChan(node)
	if nc != nil {
		nc <- channel.BulidOnTimeMessage(nil)
	}
}

func (t *CronNode) HandleMessage(msg interface{}) {

}

func (t *CronNode) GetMsgHandler() MsgHandler {
	return t
}

func (t *CronNode) Stop() {
	t.cronInstance.Stop()
}

func (t *CronNode) Name() string {
	return reflect.TypeOf(t).Elem().Name()
}
