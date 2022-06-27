package utils

import (
	"fmt"
	"testing"
)

/**
 *
 * @author jensen.chen
 * @date 2022/6/27
 */
func TestProtolGet(t *testing.T) {
	pt, _ := protocolFromHttp("https://127.0.0.1:8086/nacos")
	if pt != nil {
		fmt.Println(pt.Ip)
		fmt.Println(pt.Context)
		fmt.Println(pt.Port)
		fmt.Println(pt.Schema)
	}
}
