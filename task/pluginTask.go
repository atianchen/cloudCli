package task

import (
		"cloudCli/plugin/ctx"
		"cloudCli/plugin"
		"cloudCli/cfg"
		"github.com/robfig/cron/v3"
		"reflect"
		"log"
     )
/**
 * 负责管理所有plugin
 */
var cronInstance = cron.New()
type PluginTask struct{
	AbstractTask
	PluginList []plugin.Plugin
}

func (t *PluginTask) Start(params TaskParams){ 
   /**
    * 从配置脚本加载，包括插件配置、定时的配置等
    * 需要根据 params的内容，来决定执行那些PLUGIN
    */
   ctx := ctx.CreateContext()
   pluginParams := plugin.ExecuteParams{}
   cron := cfg.GetConfig("cli.inspect.cron")
   if (cron!=nil){
   	log.Println("Cron ",cron)
		_,err := cronInstance.AddFunc("0/10 * * * *", func(){
			for _,instance := range t.PluginList{
				instance.Execute(ctx,pluginParams)
		    }
		 })  
		if (err!=nil){
			log.Println(err)
		}
	}
	cronInstance.Start()
}

func (t *PluginTask) Stop(){
   cronInstance.Stop()
}

func  (t *PluginTask) Name()string{
    return reflect.TypeOf(t).Elem().Name()
}

func initPluginTask()Task{
	inspectPlugin := plugin.InspectPlugin{}
	task := PluginTask{PluginList:[]plugin.Plugin{&inspectPlugin}}
	return &task
}
