package plugin

import (
	"cloudCli/ctx"
	"cloudCli/utils/log"
)

type MailBoxConfig struct {
	Host     string //邮箱服务器
	Protocol string //协议 smtp imap
	User     string
	Passowrd string
}

/**
 * 发送邮件插件
 */
type MailPlugin struct {
	config MailBoxConfig
}

func (t *MailPlugin) Execute(context ctx.Context, params ExecuteParams) {
	log.Info("Execute MailPlugin")
}
