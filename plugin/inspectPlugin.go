package plugin

import (
	"cloudCli/ctx"
	"cloudCli/utils/log"
)

/**
 * 系统巡检插件
 */
type InspectPlugin struct {
	BasePlugin
}

func (t *InspectPlugin) Execute(context ctx.Context, params ExecuteParams) {
	log.Info("Execute InspectPlugin")
}
