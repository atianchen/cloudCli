package domain

import ("cloudCli/common")

type DocInfo struct{
	common.BaseObj
	Id int64
	Name string
	Path string
	ModifyDate int64 `db:"modifyDate"`//文件最后的修改日期 
	LastCheckDate int64 `db:"lastCheckDate"` //文件最后的检测日期
	Hash string //文件的HASH值得
}