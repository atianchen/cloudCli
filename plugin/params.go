package plugin

import (
		"os"
		"cloudCli/common"
	)

/**
 * 任务执行的参数
 */
type ExecuteParams struct
{
	common.ModalMap
	/**
	 * 附件
	 */
	attachs []*os.File
	
}