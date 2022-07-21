package notify

/**
 *
 * @author jensen.chen
 * @date 2022/7/21
 */
type NotifyHistoryDto struct {
	Id       string `json:"id"`
	Receiver string `json:"receiver"` //接收人
	Content  string `json:"content"`  //通知内容
	SendTime int64  `json:"sendTime"`
}
