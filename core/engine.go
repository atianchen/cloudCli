package core

import (
	"cloudCli/ctx"
	"cloudCli/node"
)

/**
 * 执行引擎
 */
type Engine interface {
	execute(ctx ctx.Context, task node.Task)
}
