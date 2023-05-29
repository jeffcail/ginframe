package email

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

// SendMail 发送邮件
// 支持 163邮箱、QQ邮箱、126邮箱
// from 发件人邮箱 `xxxx@163.com`、`xxxx@qq.com`、`xxxx@126.com`
// to 接受人邮箱
// title 邮箱标题
// content 邮箱内容
// mailServer 163邮箱服务地址 163: `smtp.163.com`、 QQ邮箱: `smtp.qq.com`、 126邮箱: `smtp.126.com`
// mailServerPort 163邮箱端口 163: `:465` 、 QQ邮箱端口: `:465`、 126邮箱: `:465`
// mailPassword 163邮箱、QQ邮箱、126邮箱 授权码 `xxxxxxxxxxx`
func SendMail(from, to, title, content, mailServer, mailServerPort, mailPassword string) (err error) {
	e := email.NewEmail()
	e.From = from
	e.To = []string{to}
	e.Subject = title
	e.HTML = []byte(content)
	return e.SendWithTLS(fmt.Sprintf("%s%s", mailServer, mailServerPort),
		smtp.PlainAuth("", from, mailPassword, mailServer),
		&tls.Config{InsecureSkipVerify: true, ServerName: mailServer})
}

// SendGmailEmail 发送谷歌邮件
// 支持 Gmail邮箱
// from 发件人邮箱 `xxxx@gmail.com`
// to 接受人邮箱
// title 邮箱标题
// content 邮箱内容
// mailServer Gmail邮箱服务地址 `smtp.gmail.com`
// mailServerPort Gmail邮箱端口 `:587`
// mailPassword Gmail邮箱登陆密码
func SendGmailEmail(from, to, title, content, mailServer, mailServerPort, mailPassword string) (err error) {
	e := email.NewEmail()
	e.From = from
	e.To = []string{to}
	e.Subject = title
	e.HTML = []byte(content)
	return e.SendWithStartTLS(fmt.Sprintf("%s%s", mailServer, mailServerPort),
		smtp.PlainAuth("", from, mailPassword, mailServer),
		&tls.Config{InsecureSkipVerify: true, ServerName: mailServer})
}

// SendGmail  发送谷歌邮件 支持同时给多人发送
func SendGmail(from string, to []string, content []byte, mailServer, mailServerPort, mailPassword string) (err error) {
	auth := smtp.PlainAuth("", from, mailPassword, mailServer)
	return smtp.SendMail(mailServer+mailServerPort, auth, from, to, content)
}
