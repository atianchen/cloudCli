package ctx

/**
 *
 * @author jensen.chen
 * @date 2022/7/26
 */
var APPINFO = appInfo{}

type appInfo struct {
	APP_NAME    string
	SERVER_BIND string
	SERVER_PORT uint
}

func InitAppInfo(name string, bind string, port uint) {
	APPINFO.APP_NAME = name
	APPINFO.SERVER_BIND = bind
	APPINFO.SERVER_PORT = port
}
