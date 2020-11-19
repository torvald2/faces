package adaptors

import (
	"io"
	"strings"

	"atbmarket.comfaceapp/config"
	"github.com/jordan-wright/email"
)

func SendReport(attatch io.Reader, attachName string) error {
	conf := config.GetConfig()
	mail := email.NewEmail()
	mail.Subject = "Ежедневная рассылка Приход/Уход"
	mail.From = "Эксперимент уход/приход"
	mail.To = strings.Split(conf.Emails, ",")
	mail.Text = []byte("Ежедневный отчет по приходу уходу сотрудников во вложении")
	mail.Attach(attatch, attachName, "")

}
