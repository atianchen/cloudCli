package cfg

import (
	"cloudCli/common"
	"cloudCli/utils"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"reflect"
	"strings"
)

/**
 * 系统配置
 * @author jensen.chen
 * @date 2022-05-21
 */
type Config struct {
	Data map[string]interface{}
}

var CliConfig Config = Config{}

/**
 * 获取配置
 * 配置的KEY可以用.分割，比如inspect.cron\
 * 如果找不到配置,则返回指定的默认值
 */
func GetConfigValue(key string, defaultValue interface{}) interface{} {
	val := GetConfig(key)
	if val != nil {
		return val
	} else {
		return defaultValue
	}
}

/**
 * 获取配置
 * 配置的KEY可以用.分割，比如inspect.cron
 */
func GetConfig(key string) interface{} {
	keyAry := strings.Split(key, common.KEY_DELIMITER)
	var val interface{} = CliConfig.Data
	for _, itemKey := range keyAry {
		val = getMapValue(val.(map[string]interface{}), itemKey)
		if val == nil {
			break
		}

	}
	return val
}

func GetKeys(data map[string]interface{}) []string {
	j := 0
	keys := make([]string, len(data))
	for k := range data {
		keys[j] = k
		j++
	}
	return keys
}

func getMapValue(data map[string]interface{}, key string) interface{} {
	rs, _ := data[key]
	return rs
}

/**
配置项自动写入struct变量
@key 配置项的ROOT键值
@target Struct指针
*/
func ConfigMapping(key string, target interface{}) {
	config := GetConfig(key)
	if config != nil && reflect.ValueOf(config).Kind() == reflect.Map {
		data := config.(map[string]interface{})
		utils.MapToStruct(data, target, "config")
	}
}

func Load(path string) {
	log.Println("Read Config", path)
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	decode := yaml.NewDecoder(f)
	decode.Decode(&(CliConfig.Data))
	log.Println(CliConfig.Data)
}
