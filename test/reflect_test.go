package test

import (
	"cloudCli/driver"
	"fmt"
	"reflect"
	"testing"
)

/**
 *反射测试
 * @author jensen.chen
 * @date 2022/6/28
 */
func TestReflect(t *testing.T) {
	cfg := &driver.NacosConfig{}
	ct := reflect.TypeOf(cfg).Elem()
	vt := reflect.ValueOf(cfg).Elem()
	vt.Field(1).SetString("11111")
	for i := 0; i < ct.NumField(); i++ {
		if ct.Field(i).Type.Kind() == reflect.Pointer {
			fmt.Println(ct.Field(i).Type.Elem())
		} else {
			fmt.Println(vt.Field(i).String())
			fmt.Println(ct.Field(i).Name)
		}

	}
}
