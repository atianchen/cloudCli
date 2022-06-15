package node

import (
	_ "github.com/robfig/cron/v3"
	"reflect"
)

/**
定时任务调度
*/
type CronNode struct {
	AbstractNode
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

}

func (t *CronNode) Stop() {

}

func (t *CronNode) Name() string {
	return reflect.TypeOf(t).Elem().Name()
}
