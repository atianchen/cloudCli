package driver

import (
	"cloudCli/utils"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
)

/**
 * Nacos Client
 * @author jensen.chen
 * @date 2022/6/27
 */
type NacosConfig struct {
	Server    utils.Protocol //服务器信息
	User      string         //用户名
	Password  string         //密码
	LogDir    string         //日志目录
	CacheDir  string         //缓存目录
	NameSpace string         //命名空间
}
type NacosClient struct {
	config       *NacosConfig
	configClient config_client.IConfigClient
}

/**
根据配置文件创建NacosClient
*/
func createNacosClientFromConfig() (*NacosClient, error) {
	config.getConfg
}

func CreateNacosClient(config *NacosConfig) (*NacosClient, error) {
	client := &NacosClient{config: config}
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
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		return nil, err
	}
	client.configClient = configClient
	return client, nil
}

func (*NacosClient) Create(config *NacosConfig) {

}

func (*NacosClient) Release() {

}
