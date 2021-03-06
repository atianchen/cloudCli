package domain

import "cloudCli/common"

const (
	DOCHIS_STATUS_PENDING = iota //待处理
	DOCHIS_STATUS_END            //已处理
	DOCHIS_STATUS_EXPIRED        //过期
)
const (
	DOCHIS_RESULT_RESERVE = 1 //保留变更
	DOCHIS_RESULT_REVERT  = 2 //还原变更
)

/**
 * 文件变更历史
 * @author jensen.chen
 * @date 2022/7/5
 */
type DocHistory struct {
	common.BaseObj
	Id           string
	DocId        string `db:"doc_id"`
	Name         string
	Path         string
	NestedPath   string `db:"nested_path"`
	ModifyTime   int64  `db:"modify_time"` //变更时间
	Hash         string
	Raw          string //原始的文件内容
	Content      string //变更后的文件让内容
	Status       int    //状态
	HandleResult int    `db:"handle_result"`
	HandleTime   int64  `db:"handle_time"` //处理时间
	Handler      string
	Opinion      string
}
