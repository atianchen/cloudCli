package sys

/**
 *
 * @author jensen.chen
 * @date 2022/7/11
 */
type ParamDto struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Code  string `json:"code"`
	Val   string `json:"val"`
	Group string `json:"group"` //分组
}
