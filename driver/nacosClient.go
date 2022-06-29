package driver

import (
	"cloudCli/cfg"
	"cloudCli/utils"
	"errors"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"strconv"
	"time"
)

/**
 * Nacos Client
 * @author jensen.chen
 * @date 2022/6/27
 */
var clientConfig *constant.ClientConfig
var serverConfig []constant.ServerConfig
var nacosClient *NacosClient

/**
Nacos配置
*/
type NacosConfig struct {
	Server    *utils.Protocol //服务器信息
	User      string          //用户名
	Password  string          //密码
	LogDir    string          `config:"logDir"`    //日志目录
	CacheDir  string          `config:"cacheDir"`  //缓存目录
	NameSpace string          `config:"nameSpace"` //命名空间
}

/**
Nacos Client
*/
type NacosClient struct {
	configClient config_client.IConfigClient
	namingClient naming_client.INamingClient
}

/**
根据配置文件创建NacosClient
*/
func CreateNacosClientFromConfig() (*NacosClient, error) {
	/**
	读取Nacos配置
	*/
	nacosCfg := NacosConfig{}
	cfg.ConfigMapping("cli.nacos", &nacosCfg)
	addr := cfg.GetConfig("cli.nacos.addr")
	if addr != nil {
		protocol, err := utils.ProtocolFromHttp(addr.(string))
		if err != nil {
			return nil, err
		}
		nacosCfg.Server = protocol
		return CreateNacosClient(&nacosCfg)
	} else {
		return nil, errors.New("InValid Nacos Addr")
	}
}

/**
创建NacosClient
*/
func CreateNacosClient(config *NacosConfig) (*NacosClient, error) {
	nacosClient := &NacosClient{}
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(

			config.Server.Ip,
			config.Server.Port,
			constant.WithScheme(config.Server.Schema),
			constant.WithContextPath(config.Server.Context),
		),
	}
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(""), //When namespace is public, fill in the blank string here.
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(config.LogDir),
		constant.WithCacheDir(config.CacheDir),
		constant.WithLogLevel("debug"),
	)
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		return nil, err
	}
	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		return nil, err
	}
	nacosClient.namingClient = namingClient
	nacosClient.configClient = configClient
	return nacosClient, nil
}

/**
发布配置项
*/
func (inst *NacosClient) PublishConfig(dataId string, group string, content string) (bool, error) {
	return inst.configClient.PublishConfig(vo.ConfigParam{
		DataId:  dataId,
		Group:   group,
		Content: content})
}

/**
获取配置项
*/
func (inst *NacosClient) GetConfig(dataId string, group string) (string, error) {
	return inst.configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
}

/**
删除配置项
*/
func (inst *NacosClient) DeleteConfig(dataId string, group string) (bool, error) {
	return inst.configClient.DeleteConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
}

/**
注册服务
*/
func (inst *NacosClient) RegisteInstance(instance ServiceInstance) (bool, error) {
	/*Ip:          "10.0.0.11",
	Port:        8848,
		ServiceName: "demo.go",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc":"shanghai"},
		ClusterName: "cluster-a", // default value is DEFAULT
		GroupName:   "group-a",*/
	param := vo.RegisterInstanceParam{
		Ip:        instance.Ip,
		Weight:    1,
		Port:      instance.Port,
		Ephemeral: true,
	}
	if len(instance.Group) > 0 {
		param.GroupName = instance.Group
	}
	if len(instance.Cluster) > 0 {
		param.ClusterName = instance.Cluster
	}
	metaData := make(map[string]string)
	metaData["ts"] = strconv.Itoa(int(time.Now().UnixNano() / 1e6))
	if instance.Data != nil {
		for i, v := range instance.Data {
			metaData[i] = v
		}
	}
	param.Metadata = metaData
	return inst.namingClient.RegisterInstance(param)
}

/**
服务取消注册
*/
func (inst *NacosClient) DeRegisterInstance(instance ServiceInstance) {
	param := vo.DeregisterInstanceParam{
		Ip:        instance.Ip,
		Port:      instance.Port,
		Ephemeral: true,
	}
	if len(instance.Group) > 0 {
		param.GroupName = instance.Group
	}
	if len(instance.Cluster) > 0 {
		param.Cluster = instance.Cluster
	}
	inst.namingClient.DeregisterInstance(param)
}

/**
获取服务信息
*/
func (inst *NacosClient) GetRegisterInstance(name string, group string, clusters []string) ([]ServiceInstance, error) {
	param := vo.GetServiceParam{
		ServiceName: name,
	}
	if len(group) > 0 {
		param.GroupName = group
	}
	if len(clusters) > 0 {
		param.Clusters = clusters
	}
	result, err := inst.namingClient.GetService(param)
	if err != nil {
		return nil, err
	}
	instances := make([]ServiceInstance, len(result.Hosts))
	for index, inst := range result.Hosts {
		instances[index] = ServiceInstance{
			Ip:          inst.Ip,
			Port:        inst.Port,
			Cluster:     inst.ClusterName,
			ServiceName: inst.ServiceName,
			Data:        inst.Metadata,
		}
	}
	return instances, nil
}
