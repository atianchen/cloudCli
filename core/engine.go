package core

import (
        "cloudCli/plugin/ctx"
        "cloudCli/task"
        )
/**
 * 执行引擎
 */
type Engine  interface
{
	execute (ctx ctx.Context,task task.Task)
}