package node

import (
	"cloudCli/cfg"
	"cloudCli/channel"
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

					instance.Execute(ctx.CreateContext(instance), pluginParams)
				}
			})
			if err != nil {
				log.Info(err)
			}
		}*/
}

func (t *CronNode) Start() {
	t.cronInstance = cron.New()
	cronExpress := cfg.GetConfig("cli.timer")
	if cronExpress != nil {
		for node, express := range cronExpress.(map[string]interface{}) {
			t.cronInstance.AddFunc(express.(string), func() {
				/**
				定时任务通知
				*/
				nc := channel.GetChan(node)
				if nc != nil {
					nc <- channel.BulidOnTimeMessage(nil)
				}
			})
		}
	}
}

func (t *CronNode) Stop() {
	t.cronInstance.Stop()
}

func (t *CronNode) Name() string {
	return reflect.TypeOf(t).Elem().Name()
}
