package plugin

import  (
        "cloudCli/plugin/ctx"
        "log"
    )
/**
 * 系统巡检插件
 */
type InspectPlugin struct{
    
}

func (t *InspectPlugin) Execute(context ctx.Context,params ExecuteParams){
    log.Println("Execute InspectPlugin")
}