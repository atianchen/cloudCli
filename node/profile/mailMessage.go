package profile

import (
	"cloudCli/component/nofity"
	"cloudCli/domain"
	"cloudCli/repository"
	"cloudCli/utils/tpl"
	"errors"
	"strconv"
)

/**
 * 预警信息发送
 * @author jensen.chen
 * @date 2022/7/12
 */
func SendMailAlarm(dh *domain.DocHistory) error {
	repo := repository.ParamRepository{}
	var params []domain.Param
	/**
	获取邮箱参数
	*/
	err := repo.GetGroupParams(&params, domain.PARAM_GROUP_MAIL)
	if err != nil {
		return err
	}
	var mailServer nofity.MailHost
	for _, param := range params {
		switch param.Code {
		case domain.PARAM_MAIL_HOST:
			{
				mailServer.Host = param.Val
			}
		case domain.PARAM_MAIL_ADDR:
			{
				mailServer.Addr = param.Val
			}
		case domain.PARAM_MAIL_PORT:
			{
				intVal, err := strconv.Atoi(param.Val)
				if err == nil {
					mailServer.Port = intVal
				} else {
					return err
				}
			}
		case domain.PARAM_MAIL_USER:
			{
				mailServer.User = param.Val
			}
		case domain.PARAM_MAIL_PWD:
			{
				mailServer.Pwd = param.Val
			}
		}
	}
	if len(mailServer.Host) < 1 || len(mailServer.Addr) < 1 || len(mailServer.User) < 1 || len(mailServer.Pwd) < 1 {
		return errors.New("Missing Parameter")
	}
	/**
	构建邮件内容
	*/
	tplParam, err := repo.GetGroupSpecParam(domain.PARAM_GROUP_PROFILE, domain.PARAM_PROFILE_MAILTEMPLATE)
	if err != nil {
		return err
	}
	mailContent, err := tpl.ProcessTemplate(tplParam.Val, dh)
	if err != nil {
		return err
	}
	var mailItem nofity.MailItem
	mailItem.Content = mailContent
	recvParam, err := repo.GetGroupSpecParam(domain.PARAM_GROUP_PROFILE, domain.PARAM_PROFILE_MAILRECEIVER)
	if err != nil {
		return err
	}
	mailItem.To = recvParam.Val
	return nofity.SendMail(&mailServer, &mailItem)
}
