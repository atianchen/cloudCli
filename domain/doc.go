package domain

import (
	"cloudCli/common"
)

const (
	DOC_TYPE_DISKFILE = iota //磁盘文件
	DOC_TYPE_JARFILE         //jar文件
	DOC_TYPE_FTP
)

type DocInfo struct {
	common.BaseObj
	Id         string
	Name       string
	Path       string
	NestedPath string `db:"nested_path"`
	Type       int8
	Content    string
	CreateTime int64  `db:"create_time"`
	ModifyTime int64  `db:"modify_time"` //文件最后的修改日期
	CheckTime  int64  `db:"check_time"`  //文件最后的检测日期
	Hash       string //文件的HASH值得
}
