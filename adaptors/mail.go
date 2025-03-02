package adaptors

import (
	"fmt"
	"io"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
	"github.com/torvald2/faces/config"
)

func SendReport(attatch io.Reader, attachName string, emails string) error {
	content_type := "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	conf := config.GetConfig()
	mail := email.NewEmail()
	mail.Subject = "Ежедневная рассылка Приход/Уход"
	mail.From = fmt.Sprintf("Эксперимент уход/приход <%v>", conf.EMailUser)
	mail.To = strings.Split(emails, ",")
	mail.Text = []byte("Ежедневный отчет по приходу уходу сотрудников во вложении")
	if _, err := mail.Attach(attatch, attachName, content_type); err != nil {
		return err
	}
	auth := smtp.PlainAuth("", conf.EMailUser, conf.EmailPassword, conf.EMailDomain)

	if err := mail.Send(fmt.Sprintf("%v:%v", conf.EMailDomain, conf.EMailPort), auth); err != nil {
		return err
	}
	return nil

}
