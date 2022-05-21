package cfg

import (
	"cloudCli/common"
	"strings"
	"log"
	"os"
	"gopkg.in/yaml.v3"
	)



/**
 * 系统配置
 * @author jensen.chen
 * @date 2022-05-21
 */
type Config struct{

	Data map[string]interface{}
}

var CliConfig Config = Config{}

/**
 * 获取配置
 * 配置的KEY可以用.分割，比如inspect.cron
 */
func  GetConfig(key string)interface{}{
	keyAry := strings.Split(key,common.KEY_DELIMITER)
	var val interface{} = CliConfig.Data
	for _,itemKey :=range keyAry{
	    val = getMapValue(val.(map[string]interface{}),itemKey)
	    if (val==nil){
	    	break
	    }
	    
	}
	return val
}

func getMapValue(data map[string]interface{},key string)interface{}{
	rs,_:= data[key]
	return rs
}


func Load(path string){
	log.Println("Read Config",path)
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	decode:=yaml.NewDecoder(f)
	decode.Decode(&(CliConfig.Data))
	log.Println(CliConfig.Data)
}