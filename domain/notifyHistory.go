package domain

/**
 *报警历史
 * @author jensen.chen
 * @date 2022/7/20
 */
type NotifyHistory struct {
	Id       string
	Receiver string //接收人
	Content  string //通知内容
	SendTime int64  `db:"send_time"`
}
