package controller

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"net/smtp"
	"spectrum/common/env"
	"spectrum/service/email/model"
)

func send(emailContext *model.EmailContext) error {
	msg := []byte(
		fmt.Sprintf("From: %s\r\nSubject: %s\r\nContent-Type: %s\r\n\r\n%s",
			emailContext.EmailAddrOfSender,
			emailContext.Subject,
			emailContext.GetContentTypeString(),
			emailContext.Body,
		),
	)
	if err := smtp.SendMail(
		fmt.Sprintf("%s:%d",
			env.Conf.EmailClient.SmtpAddr,
			env.Conf.EmailClient.SmtpPort,
		),
		smtp.PlainAuth("",
			env.Conf.EmailClient.EmailAddr,
			env.Conf.EmailClient.AuthCode,
			env.Conf.EmailClient.SmtpAddr,
		),
		env.Conf.EmailClient.EmailAddr,
		emailContext.EmailAddrOfReceivers,
		msg,
	); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
