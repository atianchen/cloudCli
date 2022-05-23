package task

import (
	"cloudCli/db"
	"cloudCli/cfg"
	"reflect"
	"os"
)


/**
 * 数据库实例管理
 * 负责建立、销毁数据库实例
 */
type DbManager struct{
	AbstractTask
}

func (d *DbManager) Init(){
	db.MapDb = d.initNoSqlDb()
	db.Db = d.initSqlDb()
}	

/**
 * 初始化NOSQL数据库
 */
func (d *DbManager) initNoSqlDb() db.NoSqlDb{
	dbConfig := db.FileDbConfig{}
	dir,_ := os.Getwd()
	dbConfig.DbFile = dir+"/db"
	cfgVal:=cfg.GetConfig("cli.db.badger.path")
	if (cfgVal!=nil){
		dbConfig.DbFile = cfgVal.(string)
	}
	db := &db.BadgerDb{}
	db.Connect(dbConfig) 
	return db
}

/**
 * 初始化SQLlITE3数据库链接
 */
func (d *DbManager) initSqlDb()db.SqlDb{
	dbConfig:= db.FileDbConfig{}
	dir,_ := os.Getwd()
	dbConfig.DbFile=dir+"/db/cli.db"
	cfgVal := cfg.GetConfig("cli.db.sqllite3.path")
	if (cfgVal!=nil){
		dbConfig.DbFile = cfgVal.(string)
	}
	db := &db.SqlLiteDb{}
	db.Connect(dbConfig) 
	return db
}
	
func (d *DbManager) Start(params TaskParams){

}


func (d *DbManager) Stop(){
	if (db.MapDb!=nil){
		db.MapDb.Release()
	}
}

func  (t *DbManager) Name()string{
    return reflect.TypeOf(t).Elem().Name()
}


