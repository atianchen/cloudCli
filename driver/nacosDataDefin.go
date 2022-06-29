package driver

/**
 * Nacos Service的数据结构定义
 * @author jensen.chen
 * @date 2022/6/29
 */
type ServiceInstance struct {
	Ip          string
	Port        uint64
	ServiceName string
	Cluster     string
	Group       string
	Data        map[string]string //附加数据
}
