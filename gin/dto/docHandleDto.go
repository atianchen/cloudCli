package dto

/**
 *
 * @author jensen.chen
 * @date 2022/7/19
 */
type DocHandleDto struct {
	Id           string
	HandleResult int   `db:"handle_result"`
	HandleTime   int64 `db:"handle_time"` //处理时间
	Handler      string
	Opinion      string
	Status       int
}
