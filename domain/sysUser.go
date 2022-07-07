package domain

import "cloudCli/common"

/**
 *
 * @author jensen.chen
 * @date 2022/7/7
 */

const (
	SYSUSER_STATUS_DISABLE = iota
	SYSUSER_STATUS_VALID
)

/**
 * 系统用户
 * @author jensen.chen
 * @date 2022/7/5
 */
type SysUser struct {
	common.BaseObj
	Id     string
	Code   string
	RoleId string `db:"role_id"`
	Name   string
	Pwd    string
	Status int
}
