package domain

/**
内置分组
*/
const (
	PARAM_GROUP_MAIL    = "mail"
	PARAM_GROUP_PROFILE = "profile"
)

/**
分组的配置项
*/
/**
邮件配置项
*/
const (
	PARAM_MAIL_HOST = "host"
	PARAM_MAIL_PORT = "port"
	PARAM_MAIL_USER = "user"
	PARAM_MAIL_PWD  = "pwd"
	PARAM_MAIL_ADDR = "addr"
)

/**
文件巡检配置项
*/
const (
	PARAM_PROFILE_MAILTEMPLATE = "mail_template" //报警邮件模板
	PARAM_PROFILE_MAILRECEIVER = "mail_receiver" //报警邮件接收人
)

/**
 * 系统参数
 * @author jensen.chen
 * @date 2022/7/8
 */
type Param struct {
	Id    string
	Name  string
	Code  string
	Val   string
	Group string `db:"param_group"` //分组
}
