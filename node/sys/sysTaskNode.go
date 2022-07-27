package sys

import (
	"bytes"
	"cloudCli/cfg"
	channel2 "cloudCli/channel"
	"cloudCli/ctx"
	"cloudCli/node"
	"cloudCli/node/extend"
	"cloudCli/server"
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
	registeCenterUrl string //注册中心地址
	/*	nacosClient *driver.NacosClient*/
}

func (t *SysTaskNode) HandleMessage(msg interface{}, channel chan interface{}) {
	switch msg.(type) {
	case *channel2.CommandMessage:
		{
			switch msg.(*channel2.CommandMessage).Name {
			case channel2.MESSAGE_ONTIME:
				{
					if err := t.nodePing(); err != nil {
						log.Error(err.Error())
					}
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

	leaderAddr, err := cfg.GetConfig("cli.cloud.addr")
	if err != nil {
		log.Error("Leader Addr Miss")
		return err
	}
	t.registeCenterUrl = leaderAddr.(string)

	/*	utils.MapToStruct(serviceConfig.(map[string]interface{}), &t.serviceInfo, "")
		if len(t.serviceInfo.Name) < 1 {
			t.serviceInfo.Name = t.serviceInfo.Ip
		}
		t.serviceInfo.Data = make(map[string]string)*/
	return nil
}

func (t *SysTaskNode) Start(context *node.NodeContext) {
	if err := t.registeNode(); err != nil {
		log.Error(err)
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

func (t *SysTaskNode) buildPayload() ([]byte, error) {
	bytes, err := json.Marshal(server.DeployNodeDto{ctx.APPINFO.APP_NAME,
		ctx.APPINFO.SERVER_BIND,
		ctx.APPINFO.SERVER_PORT})
	if err != nil {
		return nil, err
	}
	body, err := server.Encrypt(string(bytes))
	if err != nil {
		return nil, err
	}
	return json.Marshal(server.NodePayload{
		Content: body,
	})
}

func (t *SysTaskNode) registeNode() error {
	body, err := t.buildPayload()
	if err != nil {
		return err
	}
	resp, err := http.Post(t.registeCenterUrl+"/node/registe", "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("Node Regsite Error")
	}
	return nil
}

func (t *SysTaskNode) nodePing() error {
	body, err := t.buildPayload()
	if err != nil {
		return err
	}
	resp, err := http.Post(t.registeCenterUrl+"/node/ping", "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("Node Regsite Error ")
	}
	return nil
}

func (t *SysTaskNode) Stop() {

}

func (t *SysTaskNode) Name() string {
	return "sysTask"
}
