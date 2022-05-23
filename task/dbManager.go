package task

import (
	"cloudCli/db"
	"cloudCli/cfg"
	"os"
)


type DbManager struct{
	AbstractTask
}

func (d *DbManager) Init(){
	dir,_ := os.Getwd()
	dbFile:=dir+"/db"
	cfgVal:=cfg.GetConfig("cli.db.badger.path")
	if (cfgVal!=nil){
		dbFile = cfgVal.(string)
	}
	db.MapDb = &db.BadgerDbHelper{}
	db.MapDb.Connect(dbFile) 	
}
	
func (d *DbManager) Start(params TaskParams){

}


func (d *DbManager) Stop(){
	if (db.MapDb!=nil){
		db.MapDb.Release()
	}
}
