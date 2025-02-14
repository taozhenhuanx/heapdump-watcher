package email

import (
	"gopkg.in/gomail.v2"
)

// MailConn 代表邮件连接配置信息
type MailConn struct {
	User string // 发件人
	Pass string // 发件人密码或者授权码
	Host string // 邮箱地址
	Port int    // 邮箱端口 (通常是 int)
}

// Mail 代表邮件内容和收件人信息
type Mail struct {
	Conn        MailConn // 邮件连接信息
	From        string   // 发件人地址
	Subject     string   // 邮件主题
	Body        string   // 邮件内容
	To          []string // 收件人
	Cc          []string // 抄送人
	Bcc         []string // 暗送人
	Attachments []string // 附件列表
}

// Send 方法用于发送邮件
func (m *Mail) Send() error {
	// 创建 gomail 邮件消息
	message := gomail.NewMessage()
	// 发送者
	message.SetHeader("From", m.Conn.User)
	// 接收者
	message.SetHeader("To", m.To...)
	// 抄送者
	message.SetHeader("Cc", m.Cc...)
	// 暗送人
	message.SetHeader("Bcc", m.Bcc...)
	// 邮件标题
	message.SetHeader("Subject", m.Subject)
	// 邮件正文
	message.SetBody("text/html", m.Body)

	// 创建一个新的发件人
	d := gomail.NewDialer(m.Conn.Host, m.Conn.Port, m.Conn.User, m.Conn.Pass)

	// 邮件附件
	// if len(m.Attachments) > 0 {
	// 	// 添加附件
	// 	for _, attachment := range m.Attachments {
	// 		message.Attach(attachment)
	// 	}
	// }

	// 发送邮件
	return d.DialAndSend(message)
}
