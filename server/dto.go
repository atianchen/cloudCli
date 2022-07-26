package server

/**
 * DeployNodeDto
 * @author jensen.chen
 * @date 2022/7/26
 */
type DeployNodeDto struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Ip     string `json:"ip"`
	Port   int    `json:"port"`
	Ts     int64  `json:"ts"`     //最后一次PING的时间
	Status int    `json:"status"` //状态
}
type TicketVerifyDto struct {
	Ip     string `json:"ip"`
	Port   int    `json:"port"`
	Ticket string `json:"ticket"`
}
