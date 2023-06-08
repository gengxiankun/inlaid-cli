package email

import (
	"log"
	"net/smtp"

    _ "inlaid-cli/pkg/viperloader"

    "github.com/spf13/viper"
    "github.com/jordan-wright/email"
)

func Send(subject string, reply string) {
	e := email.NewEmail()
    //设置发送方的邮箱
    e.From = viper.GetString(`email.from`)
    //设置接收方的邮箱
    e.To = viper.GetStringSlice(`email.to`)
    //设置主题
    e.Subject = subject
    //设置文件发送的内容
    e.Text = []byte(reply)
    //设置服务器相关的配置
    err := e.Send(viper.GetString(`email.smtp.address`) + viper.GetString(`email.smtp.port`), smtp.PlainAuth("", viper.GetString(`email.from`), viper.GetString(`email.smtp.key`), viper.GetString(`email.smtp.address`)))
    if err != nil {
        log.Fatal(err)
    }
}
