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
				"type" INTEGER NULL,
				"content" TEXT NULL,
			    "creator"  VARCHAR(255) NULL,
				"nested_path" VARCHAR(255) NULL,
   				PRIMARY KEY('id')
			);
		CREATE TABLE IF NOT EXISTS  "inspect_doc_his" (
			    "id"  VARCHAR(100) NOT NULL,
				"doc_id" VARCHAR(100) NOT NULL,
			    "name" VARCHAR(255) NULL,
			    "path" VARCHAR(400) NULL,
				"nested_path" VARCHAR(255) NULL,
			    "modify_time" INTEGER NULL,
				"hash" VARCHAR(1000) NULL,
			    "raw" TEXT NULL,
				"content" TEXT NULL,
				"status" INTEGER NULL,
				"handle_result"	INTEGER NULL,
				"handle_time" INTEGER NULL,
				"handler" VARCHAR(255) NULL,
				"opinion" VARCHAR(255) NULL,
			    "ts" INTEGER NULL,
			    "creator"  VARCHAR(255) NULL,
   				PRIMARY KEY('id')
			);
	CREATE TABLE IF NOT EXISTS  "sys_user" (
			    "id"  VARCHAR(100) NOT NULL,
			    "code" VARCHAR(255) NULL,
			    "name" VARCHAR(255) NULL,
				"pwd" VARCHAR(255) NULL,
				"status" INTEGER NULL,
				"role_id"  VARCHAR(100) NOT NULL,
			    "ts" INTEGER NULL,
			    "creator"  VARCHAR(255) NULL,
   				PRIMARY KEY('id')
			);
	CREATE TABLE IF NOT EXISTS  "sys_param" (
			    "id"  VARCHAR(100) NOT NULL,
			    "code" VARCHAR(255) NULL,
			    "name" VARCHAR(255) NULL,
				"val" VARCHAR(255) NULL,
				"param_group" INTEGER NULL,
   				PRIMARY KEY('id')
			);
	INSERT OR IGNORE INTO "sys_user" (id,code,name,pwd,status,role_id ,ts,creator) values (1,"admin","admin","21232f297a57a5a743894a0e4a801fc3",1,"1",0,"sys");

	INSERT OR IGNORE INTO "sys_param" (id,param_group,code,name,val) values (1,"mail","host","SMTP Server Address","smtp.qq.com");
	INSERT OR IGNORE INTO "sys_param" (id,param_group,code,name,val) values (2,"mail","port","SMTP Server Port","465");
	INSERT OR IGNORE INTO "sys_param" (id,param_group,code,name,val) values (3,"mail","user","SMTP User Name","1809618127@qq.com");
	INSERT OR IGNORE INTO "sys_param" (id,param_group,code,name,val) values (4,"mail","pwd","SMTP User Password","mixjwmndvdvnbjgj");
	INSERT OR IGNORE INTO "sys_param" (id,param_group,code,name,val) values (5,"mail","addr","Mail Address","1809618127@qq.com");

	INSERT OR IGNORE INTO "sys_param" (id,param_group,code,name,val) values (6,"profile","mail_template","Alarm Mail Template","Error File：{{.Name}}");
	INSERT OR IGNORE INTO "sys_param" (id,param_group,code,name,val) values (7,"profile","mail_receiver","Alarm Mail Receiver","jensen.chen@yonyou.com.hk");
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
