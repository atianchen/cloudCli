package sys

import (
	"bytes"
	"cloudCli/cfg"
	channel2 "cloudCli/channel"
	"cloudCli/ctx"
	"cloudCli/driver"
	"cloudCli/node"
	"cloudCli/node/extend"
	"cloudCli/utils/log"
	"encoding/json"
	"errors"
	"net/http"
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
	/*	nacosClient *driver.NacosClient*/
}

func (t *SysTaskNode) HandleMessage(msg interface{}, channel chan interface{}) {
	switch msg.(type) {
	case *channel2.CommandMessage:
		{
			switch msg.(*channel2.CommandMessage).Name {
			case channel2.MESSAGE_ONTIME:
				{
					/**
					  更新注册的时间戳
					*/
					/*if t.serviceInfo.Name != "" && t.nacosClient != nil {
						log.Info("Update Service RegisterInfo:" + t.serviceInfo.Name)
						t.nacosClient.UpdateRegisteInstance(t.serviceInfo)
					}*/

				}
			}
		}
	}
}

func (d *SysTaskNode) GetMsgHandler() extend.MsgHandler {
	return d
}

func (t *SysTaskNode) Init() error {
	_, err := cfg.GetConfig("cli.cloud")
	if err != nil {
		log.Error("Skip Service Registe")
		return err
	}
	t.serviceInfo.Ip = ctx.SERVER_BIND
	t.serviceInfo.Port = uint64(ctx.SERVER_PORT)
	t.serviceInfo.Name = ctx.APP_NAME

	/*	utils.MapToStruct(serviceConfig.(map[string]interface{}), &t.serviceInfo, "")
		if len(t.serviceInfo.Name) < 1 {
			t.serviceInfo.Name = t.serviceInfo.Ip
		}
		t.serviceInfo.Data = make(map[string]string)*/
	return nil
}

func (t *SysTaskNode) registeNode() error {
	bytesContent, err := json.Marshal(t.serviceInfo)
	if err != nil {
		return err
	}
	resp, err := http.Post("http://127.0.0.1:9090/cloud/node/registe", "application/json", bytes.NewReader(bytesContent))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("Node Regsite Error ")
	}
	return nil
}

func (t *SysTaskNode) Start(context *node.NodeContext) {
	if err := t.registeNode(); err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("Node Registe Success")
	}
	/*	nc, err := driver.CreateNacosClientFromConfig()
		if err != nil {
			log.Error(err.Error())
		} else {
			t.nacosClient = nc
			_, err := t.nacosClient.RegisteInstance(t.serviceInfo)
			if err != nil {
				log.Error("Register Service Error :" + err.Error())
			}
		}*/

}

func (t *SysTaskNode) Stop() {

}

func (t *SysTaskNode) Name() string {
	return "sysTask"
}
