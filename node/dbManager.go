package node

import (
	"cloudCli/cfg"
	"cloudCli/db"
	"cloudCli/utils"
	"cloudCli/utils/log"
	"os"
	"reflect"
)

/*	Id int64
	Name string
	Path string
	ModifyDate int64//文件最后的修改日期
	LastCheckDate int64 //文件最后的检测日期
	Hash string //文件的HASH值得*/
const tableCreateSql string = `
		CREATE TABLE IF NOT EXISTS  "inspect_doc" (
			    "id" VARCHAR(100) NOT NULL,
			    "name" VARCHAR(255) NULL,
			    "path" VARCHAR(400) NULL,
				"create_time" INTEGER NULL,
			    "modify_time" INTEGER NULL,
			    "check_time" INTEGER NULL,
			    "hash" VARCHAR(1000) NULL,
			    "ts" INTEGER NULL,
			    "creator"  VARCHAR(255) NULL,
   				PRIMARY KEY('id')
			);
		CREATE TABLE IF NOT EXISTS  "inspect_doc_history" (
			    "id"  VARCHAR(100) NOT NULL,
			    "name" VARCHAR(255) NULL,
			    "path" VARCHAR(400) NULL,
			    "modify_time" INTEGER NULL,
			    "raw" TEXT NULL,
				"content" TEXT NULL,
				"status" INTEGER NULL,
				"handle_time" INTEGER NULL,
				"handler" VARCHAR(255) NULL,
			    "ts" INTEGER NULL,
			    "creator"  VARCHAR(255) NULL,
   				PRIMARY KEY('id')
			);
		`

/**
 * 数据库实例管理
 * 负责建立、销毁数据库实例
 */
type DbManager struct {
	AbstractNode
}

func (d *DbManager) HandleMessage(msg interface{}) {

}

func (d *DbManager) GetMsgHandler() MsgHandler {
	return d
}

func (d *DbManager) Init() {
	db.MapDbInst = d.initNoSqlDb()
	db.DbInst = d.initSqlDb()
	_, err := db.DbInst.Execute(tableCreateSql)
	log.Info(err)
}

/**
 * 初始化NOSQL数据库
 */
func (d *DbManager) initNoSqlDb() db.NoSqlDb {
	dbConfig := db.FileDbConfig{}
	dir, _ := os.Getwd()
	dbConfig.DbFile = dir + utils.SysSeparator() + "data"
	cfgVal, _ := cfg.GetConfig("cli.db.badger.path")
	if cfgVal != nil {
		dbConfig.DbFile = cfgVal.(string)
	}
	utils.CreateDirectory(dbConfig.DbFile)
	db := &db.BadgerDb{}
	db.Connect(dbConfig)
	return db
}

/**
 * 初始化SQLlITE3数据库链接
 */
func (d *DbManager) initSqlDb() db.SqlDb {
	dbConfig := db.FileDbConfig{}
	dir, _ := os.Getwd()
	sep := utils.SysSeparator()
	dbConfig.DbFile = dir + sep + "data" + sep + "cli.db"
	cfgVal, _ := cfg.GetConfig("cli.db.sqllite3.path")
	if cfgVal != nil {
		dbConfig.DbFile = cfgVal.(string)
	}
	utils.CreateFileDirectory(dbConfig.DbFile)
	log.Info(dbConfig.DbFile)
	db := &db.SqlLiteDb{}
	db.Connect(dbConfig)
	return db

}

func (d *DbManager) Start(context *NodeContext) {

}

func (d *DbManager) Stop() {
	if db.MapDbInst != nil {
		db.MapDbInst.Release()
	}
	if db.DbInst != nil {
		db.DbInst.Release()
	}
}

func (t *DbManager) Name() string {
	return reflect.TypeOf(t).Elem().Name()
}
