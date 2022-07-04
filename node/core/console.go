package core

import (
	"cloudCli/cfg"
	"cloudCli/channel"
	"cloudCli/node"
	"cloudCli/utils/log"
	"reflect"
	"strings"
)

/**
 * 系统控制台，用于系统的一些初始化和固有TASK执行
 * 控制台是系统的根TASK
 */
var preSetTasks = map[string]reflect.Type{"CronNode": reflect.TypeOf(node.CronNode{})} //预置任务

var sysTasks = []node.Node{&node.DbManager{}, &node.Gin{}, &SysTaskNode{}}

type Console struct {
	node.AbstractNode
}

func (c *Console) Init() {
	/**
	 * 需要根据配置决定需要执行那些系统任务
	 */
	taskConfig, _ := cfg.GetConfig("cli.node")
	if taskConfig != nil {
		taskAry := strings.Split(taskConfig.(string), ",")
		for _, taskName := range taskAry {
			task := reflect.New(preSetTasks[taskName]).Interface().(node.Node)

			sysTasks = append(sysTasks, task)
		}
	}
	/**
	 * 任务初始化
	 */
	for _, task := range sysTasks {
		task.Init()
	}
}
func (c *Console) GetMsgHandler() node.MsgHandler {
	return nil
}

func (c *Console) MessageReceive(target node.Node, channel chan interface{}) {

}

func (c *Console) Start(context *node.NodeContext) {

	for _, task := range sysTasks {
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
	for _, task := range sysTasks {
		log.Infof("Stop Task %s", task.Name())
		task.Stop()
	}

}
