package nofity

import "gopkg.in/gomail.v2"

/**
 * 邮件发送
 * @author jensen.chen
 * @date 2022/7/8
 */
type MailHost struct {
	Host string
	Port int
	User string
	Pwd  string
	Addr string //邮件地址
}

type MailItem struct {
	Subject string
	Content string
	To      string //接收方
	Cc      string //抄送
}

/**
发送邮件
*/
func SendMail(host *MailHost, mail *MailItem) error {
	m := gomail.NewMessage()
	m.SetHeader("From", host.Addr)
	m.SetHeader("To", mail.To)
	if len(mail.Cc) > 0 {
		m.SetAddressHeader("Cc", mail.Cc, "")
	}
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", mail.Content)
	d := gomail.NewDialer(host.Host, host.Port, host.User, host.Pwd)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
