package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const DEFAULT_NACOS_PORT = 80

/**
 * 协议解析
 * @author jensen.chen
 * @date 2022/6/27
 */
type Protocol struct {
	Ip      string
	Schema  string //http or https
	Port    uint64
	Context string
}

/**
从HTTP地址 获取Protocol
http://127.0.0.1:8086/nacos
*/
func ProtocolFromHttp(url string) (*Protocol, error) {
	if len(url) < 8 {
		return nil, errors.New("InValid Url")
	}
	pt := Protocol{Port: DEFAULT_NACOS_PORT}
	index := strings.Index(url, "https")
	if index == 0 {
		pt.Schema = "https"
		url = url[8:len(url)]
	} else {
		pt.Schema = "http"
		url = url[7:len(url)]
	}
	fmt.Println("url:" + url)
	ary := strings.Split(url, "/")
	/**
	获取Context Path
	*/
	if len(ary) > 1 && len(TrimBlank(ary[1])) > 0 {
		pt.Context = "/" + ary[1]
	} else {
		pt.Context = "/"
	}
	ary = strings.Split(ary[0], ":")
	if len(ary) > 1 && len(TrimBlank(ary[1])) > 0 {
		port, err := strconv.ParseInt(ary[1], 10, 64)
		if err != nil {
			return nil, err
		}
		pt.Port = uint64(port)
	}
	pt.Ip = ary[0]
	//http://127.0.0.1:8086/nacos
	return &pt, nil
}
