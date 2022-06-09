package task

import (
	"cloudCli/cfg"
	"cloudCli/ctx"
	"cloudCli/plugin"
	"cloudCli/utils/log"
	"github.com/robfig/cron/v3"
	"reflect"
)

/**
 * 负责管理所有plugin
 */
var cronInstance = cron.New()

type ScheduleTask struct {
	AbstractTask
	PluginList []plugin.Plugin
}

func (t *ScheduleTask) Init() {
	/**
	 * 从配置脚本加载，包括插件配置、定时的配置等
	 * 需要根据 params的内容，来决定执行那些PLUGIN
	 */
	t.PluginList = []plugin.Plugin{&plugin.InspectPlugin{}, &plugin.MailPlugin{}}

	pluginParams := plugin.ExecuteParams{}
	cron := cfg.GetConfig("cli.inspect.cron")
	if cron != nil {
		log.Infof("Cron %s", cron)
		_, err := cronInstance.AddFunc("* * * * *", func() {
			for _, instance := range t.PluginList {

				instance.Execute(ctx.CreateContext(), pluginParams)
			}
		})
		if err != nil {
			log.Info(err)
		}
	}
}

func (t *ScheduleTask) Start(params TaskParams) {

	cronInstance.Start()
}

func (t *ScheduleTask) Stop() {
	cronInstance.Stop()
}

func (t *ScheduleTask) Name() string {
	return reflect.TypeOf(t).Elem().Name()
}
