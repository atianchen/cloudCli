package plugin

import (
	"cloudCli/ctx"
)

/**
 * 插件
 */
type Plugin interface {
	Execute(context ctx.Context, params ExecuteParams)
}
