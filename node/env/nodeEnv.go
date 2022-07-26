package env

import (
	"cloudCli/gin"
	"cloudCli/node"
	"cloudCli/node/sys"
	"cloudCli/server"
)

/**
 * 节点环境，常量
 * @author jensen.chen
 * @date 2022/7/26
 */

var tasks = []node.Node{&node.DbManager{}, &gin.Gin{}, &sys.SysTaskNode{}, &server.ServerConsole{}}

func GetTasks() []node.Node {
	return tasks
}

func AddTask(node node.Node) {
	tasks = append(tasks, node)
}
