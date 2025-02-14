package sendAlert

import (
	"fmt"
	"heapdump_watcher/controller/sendAlert/email"
	"heapdump_watcher/setting"

	"github.com/sirupsen/logrus"
)

// 根据类型选择告警媒介, 发oss url
func SendAlertType(ossURL string) {
	switch setting.Conf.AlarmMedium.WebhookType {
	case "dingtalk":
		logrus.Println("dingtalk", "OSS URL", ossURL)
	case "email":
		logrus.Println("email", "OSS URL", ossURL)
		SenAlertEmail(ossURL)
	case "wechat":
		logrus.Println("wechat", "OSS URL", ossURL)
	default:
		logrus.Errorf("不支持该告警类型")
	}
}

// 先写死, 后面通过配置文件传递参数
func SenAlertEmail(ossURL string) {
	Body := fmt.Sprintf("<h3>JAVA 业务服务 OOM文件了, 请下载链接查看%s: </h3>", ossURL)
	// 创建邮件连接配置的实例
	mailConn := email.MailConn{
		User: "tzh971204@163.com", // 发件人
		Pass: "DJg7UxPn8FjasT3W",  // 发件人密码或者授权码
		Host: "smtp.163.com",      // 邮箱地址
		Port: 465,                 // 邮箱端口
	}

	// 创建邮件实例
	mail := email.Mail{
		Conn:        mailConn,
		From:        "tzh971204@163.com",
		To:          []string{"tzh971204@163.com"}, // 发送给用户
		Cc:          []string{"tzh971204@163.com"}, // 抄送
		Bcc:         []string{"tzh971204@163.com"}, // 暗送
		Subject:     "JAVA 业务服务 OOM文件",             // 主题
		Body:        Body,                          // 邮件正文
		Attachments: []string{""},                  // 附件
	}

	// 发送邮件
	if err := mail.Send(); err != nil {
		logrus.Error("发送邮件出错:", err)
	} else {
		logrus.Println("邮件发送成功!")
	}
}
