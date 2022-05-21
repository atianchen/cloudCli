package task

import (
    "log"
    "reflect"
    )
/**
 * 系统控制台，用于系统的一些初始化和固有TASK执行
 * 控制台是系统的根TASK
 */
var sysTasks []Task//系统任务
type Console struct{
    AbstractTask
}
func (c* Console) Start(params TaskParams){

    /**
     * 需要根据配置决定需要执行那些系统任务
     */
	sysTasks = []Task{initPluginTask()}
    for _,task :=range sysTasks{
        log.Println("Start Task ",task.Name())
        task.Start(params)
    }

}

func  (t *Console) Name()string{
    return reflect.TypeOf(t).Elem().Name()
}


func (c* Console) Stop(){
    for _,task :=range sysTasks{
        log.Println("Stop Task ",task.Name())
        task.Stop()
    }

}