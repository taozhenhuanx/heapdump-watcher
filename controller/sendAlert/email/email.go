package email

import (
	"crypto/tls"
	"fmt"
	"log"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(mailTo []string, subject string, body string) error {
	// 设置邮箱主体
	mailConn := map[string]string{
		"user": "x",         // 发件人
		"pass": "x",         // 发件人密码或者授权码
		"host": "smtp.x.cn", // 邮箱地址
		"port": "465",       // 邮箱端口
	}

	port, _ := strconv.Atoi(mailConn["port"])
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(mailConn["user"], "xx官方")) // 添加别名
	m.SetHeader("To", mailTo...)                                   // 发送给用户(可以多个)
	m.SetHeader("Cc", "******@qq.com")                             // 抄送，可以多个
	m.SetHeader("Bcc", "******@qq.com")                            // 暗送，可以多个
	m.SetHeader("Subject", subject)                                // 设置邮件主题
	m.SetBody("text/html", body)                                   // 设置邮件正文
	// m.Attach("./myIpPic.png")                                                         //附件
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"]) // 设置邮件正文
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}                               // 解决证书x509报错
	err := d.DialAndSend(m)
	return err
}
func x() {
	// 发送方
	mailTo := []string{
		"x<x@x.cn>", // 这里最好写成邮箱收发件时的这种标记格式
	}
	// 邮件主题
	subject := "Hello"
	// 邮件正文
	body := "Automatic send by Go gomail from xxx官方."
	err := SendMail(mailTo, subject, body)
	if err != nil {
		log.Print(err)
		fmt.Printf("Send fail!")
		return
	}
	fmt.Printf("send successfully!")
}
