package core

import (
	"cloudCli/cfg"
	channel2 "cloudCli/channel"
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

func (t *SysTaskNode) HandleMessage(msg interface{}) {
	switch msg.(type) {
	case *channel2.CommandMessage:
		{
			switch msg.(*channel2.CommandMessage).Name {
			case channel2.MESSAGE_ONTIME:
				{
					/**
					  更新注册的时间戳
					*/
					if t.serviceInfo.Name != "" && t.nacosClient != nil {
						log.Info("Update Service RegisterInfo:" + t.serviceInfo.Name)
						t.nacosClient.UpdateRegisteInstance(t.serviceInfo)
					}

				}
			}
		}
	}
}

func (d *SysTaskNode) GetMsgHandler() node.MsgHandler {
	return d
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

func (t *SysTaskNode) Start(context *node.NodeContext) {
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

}

func (t *SysTaskNode) Stop() {

}

func (t *SysTaskNode) Name() string {
	return "sysTask"
}
