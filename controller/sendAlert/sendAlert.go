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
func SendAlertType(ossURL, podName, nsName string) error {
	switch setting.Conf.AlarmMedium.WebhookType {
	case "dingtalk":
		return dingtalk.SendDingTalk("heapdump 告警信息,文件已经转存，请及时下载", "生产环境", ossURL, podName, nsName)
	case "email":
		return SenAlertEmail("heapdump 告警信息,文件已经转存，请及时下载", "生产环境", ossURL, podName, nsName)
	case "wechat":
		// msg, env, podName, ossURL, nsName
		return wechat.SendWeChat("heapdump 告警信息,文件已经转存，请及时下载", "生产环境", ossURL, podName, nsName)
	default:
		logrus.Errorf("不支持该告警类型")
	}
	return nil
}

// SenAlertEmail 邮件发送
func SenAlertEmail(msg, env, ossURL, podName, nsName string) error {
	body := fmt.Sprintf(
		"<h3>出现 OOM 错误, JAVA 业务服务需要您的注意。</h3>"+
			"<p>请在24小时内下载以下文件以进行分析: </p>"+
			"<p><a href=\"%s\">下载 OOM 文件</a></p>"+
			"<p>相关信息：</p>"+
			"<ul>"+
			"<li>环境: %s</li>"+
			"<li>Pod 名称: %s</li>"+
			"<li>命名空间: %s</li>"+
			"</ul>",
		ossURL, env, podName, nsName,
	)

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
		Body:        body,                         // 邮件正文
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
