package test

import (
	"cloudCli/cfg"
	"cloudCli/driver"
	"fmt"
	"testing"
)

/**
 *
 * @author jensen.chen
 * @date 2022/6/28
 */
func TestNacos(t *testing.T) {
	cfg.Load("../config.yml")
	ncfg := driver.NacosConfig{}
	cfg.ConfigMapping("cli.nacos", &ncfg)
	fmt.Println(ncfg.LogDir)
	fmt.Println(ncfg.NameSpace)
	fmt.Println(ncfg.Server.Ip)
}
