package main

import (
	"cloudCli/task"
	"cloudCli/cfg"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	)

func main() {
	fmt.Println(`                            
   ________                __   _________ 
  / ____/ /___  __  ______/ /  / ____/ (_)
 / /   / / __ \/ / / / __  /  / /   / / / 
/ /___/ / /_/ / /_/ / /_/ /  / /___/ / /  
\____/_/\____/\__,_/\__,_/   \____/_/_/  

       YonyouHK  @2022 V0.01 
  https://github.com/atianchen/cloudCli
                                         
		  `)
	sysCh := make(chan os.Signal, 1)
	signal.Notify(sysCh, syscall.SIGKILL, syscall.SIGINT)

	pwd, _ := os.Getwd()
	cfg.Load(pwd+"/config.yml")
	var rootTask task.Task = &task.Console{}
	rootTask.Start(task.TaskParams{})	
	for {
		s := <-sysCh
		switch s {
			case syscall.SIGINT:
				log.Println("Cloud Cli Exited")
				rootTask.Stop()
				return
		}
	}
}