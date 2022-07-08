package nofity

import (
	"testing"
)

func TestMailSend(t *testing.T) {
	host := MailHost{Host: "smtp.qq.com", Addr: "1809618127@qq.com", Port: 465, User: "1809618127@qq.com", Pwd: ""}
	mail := MailItem{To: "jensenchen@yonyou.com", Subject: "test", Content: "test1"}
	SendMail(&host, &mail)
}
