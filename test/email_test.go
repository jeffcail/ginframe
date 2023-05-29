package test

import (
	"fmt"
	"github.com/jeffcail/ginframe/utils/email"
	"os"
	"testing"
)

func TestSend163Mail(t *testing.T) {
	from := os.Getenv("Mail163From")
	to := "" // 接收人的邮箱
	mailPassword := os.Getenv("Mail163Pass")
	mailServer := os.Getenv("Mail163Server")
	mailServerPort := os.Getenv("Mail163ServerPort")
	title := "单元测试，邮件发送"
	content := `能发送过去吗:<h1>" + 123456 + "</h1>`
	err := email.SendMail(from, to, title, content, mailServer, mailServerPort, mailPassword)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("邮件发送成功...")
}

func TestSendQQMail(t *testing.T) {
	from := os.Getenv("MailQQFrom")
	to := "" // 接收人的邮箱
	mailPassword := os.Getenv("MailQQPass")
	mailServer := os.Getenv("MailQQServer")
	mailServerPort := os.Getenv("MailQQServerPort")
	title := "单元测试，邮件发送"
	content := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>IM注册邮件</title>
</head>
<style>
    .mail{
        margin: 0 auto;
        border-radius: 45px;
        height: 400px;
        padding: 10px;
        background-color: #CC9933;
        background: url("http://images.xxxxx.cn/note.png") no-repeat;
    }
    .code {
        color: #f6512b;
        font-weight: bold;
        font-size: 30px;
        padding: 2px;
    }
</style>
<body>
<div class="mail">
    <h3>您好:您正在测试qq邮箱发送!</h3>
    <p>下面是您的验证码:</p>
        <p class="code">%s</p>
        <p>请注意查收!谢谢</p>
</div>
<h3>如果可以请给项目点个star～<a target="_blank" href="https://github.com/jeffcail/ginframe">项目地址</a> </h3>
</body>
</html>`, "123456")
	err := email.SendMail(from, to, title, content, mailServer, mailServerPort, mailPassword)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("邮件发送成功...")
}

func TestSend126Mail(t *testing.T) {
	from := os.Getenv("Mail126From")
	to := "xxxxx@qq.com" // 接收人的邮箱
	mailPassword := os.Getenv("Mail126Pass")
	mailServer := os.Getenv("Mail126Server")
	mailServerPort := os.Getenv("Mail126ServerPort")
	title := "单元测试，邮件发送"
	content := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>IM注册邮件</title>
</head>
<style>
    .mail{
        margin: 0 auto;
        border-radius: 45px;
        height: 400px;
        padding: 10px;
        background-color: #CC9933;
        background: url("http://images.caixiaoxin.cn/note.png") no-repeat;
    }
    .code {
        color: #f6512b;
        font-weight: bold;
        font-size: 30px;
        padding: 2px;
    }
</style>
<body>
<div class="mail">
    <h3>您好:您正在测试qq邮箱发送!</h3>
    <p>下面是您的验证码:</p>
        <p class="code">%s</p>
        <p>请注意查收!谢谢</p>
</div>
<h3>如果可以请给项目点个star～<a target="_blank" href="https://github.com/jeffcail/ginframe">项目地址</a> </h3>
</body>
</html>`, "123456")
	err := email.SendMail(from, to, title, content, mailServer, mailServerPort, mailPassword)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("邮件发送成功...")
}

func TestSendEmailStartTls(t *testing.T) {
	from := os.Getenv("MailGmailFrom")
	to := "" // 接收人的邮箱
	mailPassword := os.Getenv("MailGmailPass")
	mailServer := os.Getenv("MailGmailServer")
	mailServerPort := os.Getenv("MailGmailServerPort")
	title := "单元测试，邮件发送"
	content := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>IM注册邮件</title>
</head>
<style>
    .mail{
        margin: 0 auto;
        border-radius: 45px;
        height: 400px;
        padding: 10px;
        background-color: #CC9933;
        background: url("http://images.xxxx.cn/note.png") no-repeat;
    }
    .code {
        color: #f6512b;
        font-weight: bold;
        font-size: 30px;
        padding: 2px;
    }
</style>
<body>
<div class="mail">
    <h3>您好:您正在测试qq邮箱发送!</h3>
    <p>下面是您的验证码:</p>
        <p class="code">%s</p>
        <p>请注意查收!谢谢</p>
</div>
<h3>如果可以请给项目点个star～<a target="_blank" href="https://github.com/jeffcail/ginframe">项目地址</a> </h3>
</body>
</html>`, "123456")
	err := email.SendGmailEmail(from, to, title, content, mailServer, mailServerPort, mailPassword)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("邮件发送成功...")
}

func TestSendGmail(t *testing.T) {
	from := os.Getenv("MailGmailFrom")
	to := []string{
		"", // 接收人的邮箱1
		"", // 接收人的邮箱2
	}
	mailPassword := os.Getenv("MailGmailPass")
	mailServer := os.Getenv("MailGmailServer")
	mailServerPort := os.Getenv("MailGmailServerPort")
	content := []byte(fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>IM注册邮件</title>
</head>
<style>
    .mail{
        margin: 0 auto;
        border-radius: 45px;
        height: 400px;
        padding: 10px;
        background-color: #CC9933;
        background: url("http://images.xxxx.cn/note.png") no-repeat;
    }
    .code {
        color: #f6512b;
        font-weight: bold;
        font-size: 30px;
        padding: 2px;
    }
</style>
<body>
<div class="mail">
    <h3>您好:您正在测试qq邮箱发送!</h3>
    <p>下面是您的验证码:</p>
        <p class="code">%s</p>
        <p>请注意查收!谢谢</p>
</div>
<h3>如果可以请给项目点个star～<a target="_blank" href="https://github.com/jeffcail/ginframe">项目地址</a> </h3>
</body>
</html>`, "123456"))

	err := email.SendGmail(from, to, content, mailServer, mailServerPort, mailPassword)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("邮件发送成功...")
}
