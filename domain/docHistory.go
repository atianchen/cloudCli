package domain

import "cloudCli/common"

/**
 * 文件变更历史
 * @author jensen.chen
 * @date 2022/7/5
 */
type DocHistory struct {
	common.BaseObj
	Id         string
	Name       string
	Path       string
	ModifyTime int64  `db:"modify_time"` //变更时间
	Raw        string //原始的文件内容
	Content    string //变更后的文件让内容
	Status     int    //状态
	HanleTime  int64  `db:"handle_time"` //处理时间
	Handler    string
}
