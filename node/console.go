package node

import (
	"cloudCli/cfg"
	"cloudCli/channel"
	"cloudCli/ctx"
	"cloudCli/utils/log"
	"reflect"
	"strings"
)

/**
 * 系统控制台，用于系统的一些初始化和固有TASK执行
 * 控制台是系统的根TASK
 */
var preSetTasks = map[string]reflect.Type{"plugin": reflect.TypeOf(PluginTask{})} //预置任务

var sysTasks = []Node{&DbManager{}, &Gin{}}

type Console struct {
	AbstractNode
}

func (c *Console) Init() {
	/**
	 * 需要根据配置决定需要执行那些系统任务
	 */
	taskConfig := cfg.GetConfig("cli.node")
	if taskConfig != nil {
		taskAry := strings.Split(taskConfig.(string), ",")
		for _, taskName := range taskAry {
			task := reflect.New(preSetTasks[taskName]).Interface().(Node)

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

func (c *Console) Start(context ctx.Context) {

	for _, task := range sysTasks {
		log.Infof("Start Task %s", task.Name())
		task.Start(ctx.CreateContext(task))
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
