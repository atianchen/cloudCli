package main

import (
	"cloudCli/cfg"
	"cloudCli/node"
	"cloudCli/utils/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	banner := `                            
   ________                __   _________ 
  / ____/ /___  __  ______/ /  / ____/ (_)
 / /   / / __ \/ / / / __  /  / /   / / / 
/ /___/ / /_/ / /_/ / /_/ /  / /___/ / /  
\____/_/\____/\__,_/\__,_/   \____/_/_/  

       YonyouHK  @2022 V0.01 
  https://github.com/atianchen/cloudCli
                                         
		  `

	pwd, _ := os.Getwd()
	cfg.Load(pwd + "/config.yml")
	var logger log.LogInit = &log.Log{}
	logger.Init()
	log.Info(banner)
	sysCh := make(chan os.Signal, 1)
	signal.Notify(sysCh, syscall.SIGKILL, syscall.SIGINT)

	var rootTask node.Node = &node.Console{}
	rootTask.Init()
	rootTask.Start(nil)

	for {
		s := <-sysCh
		switch s {
		case syscall.SIGINT:
			log.Info("Cloud Cli Exited")
			rootTask.Stop()
			return
		}
	}
}
