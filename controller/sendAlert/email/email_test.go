package email

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSend(t *testing.T) {
	// 创建邮件连接配置的实例
	mailConn := MailConn{
		User: "tzh971204@163.com", // 发件人
		Pass: "DJg7UxPn8FjasT3W",  // 发件人密码或者授权码
		Host: "smtp.163.com",      // 邮箱地址
		Port: 465,                 // 邮箱端口
	}

	// 创建邮件实例
	mail := Mail{
		Conn:        mailConn,
		From:        "tzh971204@163.com",
		To:          []string{"tzh971204@163.com"},         // 发送给用户
		Cc:          []string{"tzh971204@163.com"},         // 抄送
		Bcc:         []string{"tzh971204@163.com"},         // 暗送
		Subject:     "JAVA 业务服务 OOM文件",                     // 主题
		Body:        "<h3>JAVA 业务服务 OOM文件了, 请下载链接查看：</h3>", // 邮件正文
		Attachments: []string{""},                          // 附件
	}

	// 发送邮件
	if err := mail.Send(); err != nil {
		logrus.Error("发送邮件出错:", err)
	} else {
		logrus.Println("邮件发送成功!")
	}
}
