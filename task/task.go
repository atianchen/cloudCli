package task
 
import (
		"reflect"
		"cloudCli/common"
	)

type TaskParams struct{
	common.Extends
}

/**
 * 任务的相关定义
 * @author jensen.chen
 * @date 2022-05-20
 */

/**
 * 任务
 */
type Task interface{

	/**
	 * 开始任务
	 */
	Start(params TaskParams)

	/**
	 * 停止任务
	 */
	Stop()	

	/**
	 * 获取任务名称
	 */
	Name() string
}


type AbstractTask struct{

}
func  (t *AbstractTask) Name()string{
    return reflect.TypeOf(t).Elem().Name()
}


