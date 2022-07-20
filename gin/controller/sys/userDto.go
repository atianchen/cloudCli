package sys

/**
 *
 * @author jensen.chen
 * @date 2022/7/13
 */
type UserDto struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	RoleId string `json:"roleId"`
	Token  string `json:"token"`
}
