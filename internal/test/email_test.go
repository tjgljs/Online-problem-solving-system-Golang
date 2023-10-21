package test

import (
	"crypto/tls"
	"net/smtp"
	"testing"

	"getcharzp.cn/define"
	"github.com/jordan-wright/email"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <xxxx@163.com>"
	e.To = []string{"xxxxx@qq.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码：<b>123456</b>")
	// 返回 EOF 时，关闭SSL重试
	err := e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "xxxx@163.com", define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
