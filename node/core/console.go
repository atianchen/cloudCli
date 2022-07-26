package core

import (
	"cloudCli/cfg"
	"cloudCli/channel"
	"cloudCli/ctx"
	"cloudCli/node"
	"cloudCli/node/env"
	"cloudCli/node/extend"
	"cloudCli/node/profile"
	"cloudCli/utils"
	"cloudCli/utils/log"
	"reflect"
	"strings"
)

/**
 * 系统控制台，用于系统的一些初始化和固有TASK执行
 * 控制台是系统的根TASK
 */
var preSetTasks = map[string]reflect.Type{"CronNode": reflect.TypeOf(node.CronNode{}), "ProfileInspect": reflect.TypeOf(profile.ProfileInspect{})} //预置任务

type Console struct {
	node.AbstractNode
}

func (c *Console) Init() error {
	/**
	 * 需要根据配置决定需要执行那些系统任务
	 */
	defer func() {
		if r := recover(); r != nil {
			log.Error("Console Init Error", r)
		}
	}()
	taskConfig, _ := cfg.GetConfig("cli.node")
	if taskConfig != nil {
		taskAry := strings.Split(taskConfig.(string), ",")
		for _, taskName := range taskAry {
			task := reflect.New(preSetTasks[taskName]).Interface().(node.Node)
			env.AddTask(task)
		}
	}
	/**
	 * 任务初始化
	 */
	for _, task := range env.GetTasks() {
		task.Init()
	}
	c.setAppInfo()
	return nil
}

/**
设置ctx的AppInfo
*/
func (c *Console) setAppInfo() extend.MsgHandler {
	serverInfo, err := cfg.GetConfig("cli.server")
	if err != nil {
		log.Error(err.Error())
	} else {
		data := serverInfo.(map[string]interface{})
		ctx.SERVER_BIND = utils.MapValue(data, "bind", "").(string)
		ctx.SERVER_PORT = utils.MapValue(data, "port", 0).(int)
		appName, err := cfg.GetConfig("cli.cloud.name")
		if err == nil {
			ctx.APP_NAME = appName.(string)
		} else {
			log.Error(err.Error())
		}

	}
	return nil
}

func (c *Console) GetMsgHandler() extend.MsgHandler {
	return nil
}

func (c *Console) MessageReceive(target node.Node, channel chan interface{}) {

}

func (c *Console) Start(context *node.NodeContext) {

	for _, task := range env.GetTasks() {
		log.Infof("Start Task %s", task.Name())
		ntx := node.CreateNodeContext(task)
		task.Start(ntx)
		go task.MessageReceive(task, ntx.Channel)
	}
}

func (c *Console) Name() string {
	return reflect.TypeOf(c).Elem().Name()
}

func (c *Console) Stop() {
	channel.Release()
	for _, task := range env.GetTasks() {
		log.Infof("Stop Task %s", task.Name())
		task.Stop()
	}

}
