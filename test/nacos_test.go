package test

import (
	"cloudCli/cfg"
	"cloudCli/driver"
	"fmt"
	"testing"
)

func TestNacos(t *testing.T) {
	cfg.Load("../config.yml")
	client, _ := driver.CreateNacosClientFromConfig()
	client.PublishConfig("my1.yml", "my", "test")   //发布配置
	content, _ := client.GetConfig("my1.yml", "my") //获取配置
	fmt.Println(content)
	/*	ncfg := driver.NacosConfig{}
		cfg.ConfigMapping("cli.nacos", &ncfg)
		fmt.Println(ncfg.LogDir)
		fmt.Println(ncfg.NameSpace)
		fmt.Println(ncfg.Server.Ip)*/
}
