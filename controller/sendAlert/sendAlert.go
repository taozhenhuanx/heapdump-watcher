package sendAlert

import (
	"fmt"
	"heapdump_watcher/controller/sendAlert/dingtalk"
	"heapdump_watcher/controller/sendAlert/email"
	"heapdump_watcher/controller/sendAlert/wechat"
	"heapdump_watcher/setting"

	"github.com/sirupsen/logrus"
)

// SendAlertType 根据类型选择告警媒介, 发oss url
func SendAlertType(ossURL string) error {
	switch setting.Conf.AlarmMedium.WebhookType {
	case "dingtalk":
		return dingtalk.SendDingTalk("heapdump 告警信息,文件已经转存，请及时下载", "生产环境", "APP应用", ossURL)
	case "email":
		return SenAlertEmail(ossURL)
	case "wechat":
		return wechat.SendWeChat("heapdump 告警信息,文件已经转存，请及时下载", "生产环境", "APP应用", ossURL)
	default:
		logrus.Errorf("不支持该告警类型")
	}
	return nil
}

// SenAlertEmail 邮件发送
func SenAlertEmail(ossURL string) error {
	Body := fmt.Sprintf("<h3>JAVA 业务服务 OOM文件了, 请在一天内下载文件, 请下载链接查看%s: </h3>", ossURL)
	// 创建邮件连接配置的实例
	mailConn := email.MailConn{
		User: setting.Conf.AlarmEmail.User, // 发件人
		Pass: setting.Conf.AlarmEmail.Pass, // 发件人密码或者授权码
		Host: setting.Conf.AlarmEmail.Host, // 邮箱地址
		Port: setting.Conf.AlarmEmail.Port, // 邮箱端口
	}

	// 创建邮件实例
	mail := email.Mail{
		Conn:        mailConn,
		From:        setting.Conf.AlarmEmail.User, // 发件人地址
		To:          setting.Conf.AlarmEmail.To,   // 发送给用户
		Cc:          setting.Conf.AlarmEmail.Cc,   // 抄送
		Bcc:         setting.Conf.AlarmEmail.Bcc,  // 暗送
		Subject:     "JAVA 业务服务 OOM文件",            // 主题
		Body:        Body,                         // 邮件正文
		Attachments: []string{""},                 // 附件
	}

	// 发送邮件
	if err := mail.Send(); err != nil {
		logrus.Error("发送邮件出错:", err)
		return err
	} else {
		logrus.Println("邮件发送成功!")
	}
	return nil
}
