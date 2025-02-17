package email

import (
	"fmt"
	"heapdump_watcher/setting"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSend(t *testing.T) {
	// 创建邮件连接配置的实例
	mailConn := MailConn{
		User: setting.Conf.AlarmEmail.User, // 发件人
		Pass: setting.Conf.AlarmEmail.Pass, // 发件人密码或者授权码
		Host: setting.Conf.AlarmEmail.Host, // 邮箱地址
		Port: setting.Conf.AlarmEmail.Port, // 邮箱端口
	}
	Body := fmt.Sprintf("<h3>JAVA 业务服务 OOM文件了, 请下载链接查看%s: </h3>", "https://tet.com")
	// 创建邮件实例
	mail := Mail{
		Conn:        mailConn,
		From:        setting.Conf.AlarmEmail.User, // 发件人地址
		To:          setting.Conf.AlarmEmail.To,   // 发送给用户
		Cc:          setting.Conf.AlarmEmail.Cc,   // 抄送
		Bcc:         setting.Conf.AlarmEmail.Bcc,  // 暗送
		Subject:     "JAVA 业务服务 OOM文件",            // 主题
		Body:        Body,                         // 邮件正文
		Attachments: []string{""},                 // 附件
	}
	fmt.Println(mail)
	// 发送邮件
	if err := mail.Send(); err != nil {
		logrus.Error("发送邮件出错:", err)
	} else {
		logrus.Println("邮件发送成功!")
	}
}
