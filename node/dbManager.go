package node

import (
	"cloudCli/cfg"
	"cloudCli/ctx"
	"cloudCli/db"
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
			    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
			    "name" VARCHAR(255) NULL,
			    "path" VARCHAR(1000) NULL,
			    "modify_date" INTEGER NULL,
			    "last_checkdate" INTEGER NULL,
			    "hash" VARCHAR(1000) NULL,
			    "ts" INTEGER NULL,
			    "creator"  VARCHAR(255) NULL
			);
		`

/**
 * 数据库实例管理
 * 负责建立、销毁数据库实例
 */
type DbManager struct {
	AbstractNode
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
	dbConfig.DbFile = dir + "/db"
	cfgVal := cfg.GetConfig("cli.db.badger.path")
	if cfgVal != nil {
		dbConfig.DbFile = cfgVal.(string)
	}
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
	dbConfig.DbFile = dir + "/db/cli.db"
	cfgVal := cfg.GetConfig("cli.db.sqllite3.path")
	if cfgVal != nil {
		dbConfig.DbFile = cfgVal.(string)
	}
	log.Info(dbConfig.DbFile)
	db := &db.SqlLiteDb{}
	db.Connect(dbConfig)
	return db
}

func (d *DbManager) Start(context ctx.Context) {

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
