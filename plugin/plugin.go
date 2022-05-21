package plugin

import (
    "cloudCli/plugin/ctx"
)
/**
 * 插件
 */
type Plugin interface{

	 Execute(context ctx.Context,params ExecuteParams)

}


