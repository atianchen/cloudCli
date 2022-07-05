package domain

import (
	"cloudCli/common"
)

type DocInfo struct {
	common.BaseObj
	Id         string
	Name       string
	Path       string
	CreateTime int64  `db:"create_time"`
	ModifyTime int64  `db:"modify_time"` //文件最后的修改日期
	CheckTime  int64  `db:"check_time"`  //文件最后的检测日期
	Hash       string //文件的HASH值得
}
