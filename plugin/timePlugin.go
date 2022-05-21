package plugin

import (
		"net/http"
		 "cloudCli/plugin/ctx"
		 "log"
	)
/**
 * 对比系统时间与世界时钟是否差异较大
 */
type TimeInspectPlugin struct{

	/**
	 * 时间校准服务器地址
	 */
	server string
}

func (t *TimeInspectPlugin) execute(context ctx.Context,params ExecuteParams){
	/**
	 * 从远程服务器获取时间
	 */
	client := &http.Client{}
	request, _ := http.NewRequest("GET", t.server, nil)
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	request.Header.Set("Accept-Encoding", "gzip,deflate,sdch")

	response, _ := client.Do(request)
	log.Println(response)
}