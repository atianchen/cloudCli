package server

import (
	"cloudCli/db"
	"encoding/json"
	"time"
)

/**
 * 客户端信息
 * @author jensen.chen
 * @date 2022/7/25
 */
type CliClient struct {
	Ip     string
	Name   string
	Port   string
	PingTs int64 //Ping时间戳
}

const STOREKEY = "cli_clients"

type CliClientService struct {
	ClientArray []CliClient
}

func (c *CliClientService) SaveClientInfo(client CliClient) error {
	itemIndex := -1
	for index, ct := range c.ClientArray {
		if ct.Ip == client.Ip && ct.Port == client.Port {
			itemIndex = index
			break
		}
	}
	if itemIndex > -1 {
		c.ClientArray[itemIndex].PingTs = client.PingTs
		c.ClientArray[itemIndex].Ip = client.Ip
		c.ClientArray[itemIndex].Name = client.Name
		c.ClientArray[itemIndex].Port = client.Port
	} else {
		c.ClientArray = append(c.ClientArray, client)
	}
	return c.persist()
}

func (c *CliClientService) RemoveClientInfo(client CliClient) error {
	c.ClientArray = append(c.ClientArray, client)
	itemIndex := -1
	for index, ct := range c.ClientArray {
		if ct.Ip == client.Ip && ct.Port == client.Port {
			itemIndex = index
			break
		}
	}
	if itemIndex > -1 {
		c.ClientArray = append(c.ClientArray[:itemIndex], c.ClientArray[itemIndex+1])
	}
	return c.persist()
}

func (c *CliClientService) persist() error {
	content, err := json.Marshal(c.ClientArray)
	if err == nil {
		db.MapDbInst.Set(STOREKEY, string(content), time.Minute*0)
		return nil
	} else {
		return err
	}
}

func (c *CliClientService) load() error {
	content, err := db.MapDbInst.GetBytes(STOREKEY)
	if err != nil {
		return nil
	}
	data := make([]CliClient, 0)
	err = json.Unmarshal(content, &data)
	if err == nil {
		c.ClientArray = data
	}
	return err
}
