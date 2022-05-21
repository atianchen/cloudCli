package cfg

import (
	"cloudCli/common"
	"strings"
	"log"
	)


/**
 * 系统配置
 * @author jensen.chen
 * @date 2022-05-21
 */
type Config struct{

	Data map[string]interface{}
}

/**
 * 获取配置
 * 配置的KEY可以用.分割，比如inspect.timer
 */
func (c *Config) Get(key string)interface{}
{
	var keyAry := springs.Split(key,common.KEY_DELIMITER)
	var val interface{}
	for _,itemKey :=range keyAry{
	  val = getMapValue(c.Data,itemKey)
	  if val!=nil){
	  	switch (val.(type))
	  	{
	  		case map[string]interface{}:
	  			val = 
	  	}
	  } else {
	  	break
	  }
	}
	return val
}

func getMapValue(data *map[string]interface{},key string)interface{}{
	rs,_:= data[key]
	return rs
}