package core

import (
	"cloudCli/cfg"
	channel2 "cloudCli/channel"
	"cloudCli/ctx"
	"cloudCli/driver"
	"cloudCli/node"
	"cloudCli/utils"
	"cloudCli/utils/log"
)

/**
 * 系统任务
* 定时注册等
 * @author jensen.chen
 * @date 2022/6/30
*/
type SysTaskNode struct {
	node.AbstractNode
	serviceInfo driver.ServiceInstance
	nacosClient *driver.NacosClient
}

func (t *SysTaskNode) Init() {
	serviceConfig, err := cfg.GetConfig("cli.cloud")
	if err != nil {
		log.Error("Skip Service Registe")
	}
	utils.MapToStruct(serviceConfig.(map[string]interface{}), &t.serviceInfo, "")
	if len(t.serviceInfo.Name) < 1 {
		t.serviceInfo.Name = t.serviceInfo.Ip
	}
	t.serviceInfo.Data = make(map[string]string)
}

func (t *SysTaskNode) Start(context ctx.Context) {
	nc, err := driver.CreateNacosClientFromConfig()
	if err != nil {
		log.Error(err.Error())
	} else {
		t.nacosClient = nc
		_, err := t.nacosClient.RegisteInstance(t.serviceInfo)
		if err != nil {
			log.Error("Register Service Error :" + err.Error())
		}
	}
	//t.dispatch(
}

func (t *SysTaskNode) Stop() {

}

func (t *SysTaskNode) Name() string {
	return "sysTask"
}

func (t *SysTaskNode) HandleMessage(msg interface{}) *channel2.AsyncResponse {
	switch msg.(type) {
	case channel2.CommandMessage:
		{
			switch msg.(*channel2.CommandMessage).Name {
			case channel2.MESSAGE_ONTIME:
				{
					if t.serviceInfo.Name != "" && t.nacosClient != nil {
						log.Info("Update Service RegisterInfo:" + t.serviceInfo.Name)
						t.nacosClient.RegisteInstance(t.serviceInfo)
					}

				}
			}
		}
	}
	return nil
}
