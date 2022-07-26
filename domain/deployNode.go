package domain

/**
 *部署的节点
 * @author jensen.chen
 * @date 2022/7/26
 */
type DeployNode struct {
	Id     string
	Name   string
	Ip     string
	Port   int
	Ts     int64 //最后一次PING的时间
	Status int   //状态
}
