package task

import (
    "log"
    "reflect"
    "cloudCli/cfg"
    "strings"
    )
/**
 * 系统控制台，用于系统的一些初始化和固有TASK执行
 * 控制台是系统的根TASK
 */
var preSetTasks = map[string]reflect.Type{"plugin":reflect.TypeOf(PluginTask{})}//预置任务

var sysTasks []Task

type Console struct{
    AbstractTask
}

func (c* Console) Init(){
  /**
     * 需要根据配置决定需要执行那些系统任务
     */
    sysTasks = []Task{}
    taskConfig := cfg.GetConfig("cli.task")
    if (taskConfig!=nil){
        taskAry := strings.Split(taskConfig.(string),",")
        for _,taskName :=range taskAry{
            task := reflect.New(preSetTasks[taskName]).Interface().(Task)
            task.Init()
            sysTasks = append(sysTasks,task)
        }
    }
   
}

func (c* Console) Start(params TaskParams){

    for _,task :=range sysTasks{
        log.Println("Start Task ",task.Name())
        task.Start(params)
    }
}

func  (c *Console) Name()string{
    return reflect.TypeOf(c).Elem().Name()
}


func (c* Console) Stop(){
    for _,task :=range sysTasks{
        log.Println("Stop Task ",task.Name())
        task.Stop()
    }

}