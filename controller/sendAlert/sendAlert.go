package sendAlert

import (
	"fmt"
	"heapdump_watcher/setting"
)

// 根据类型选择告警媒介, 发oss url
func SendAlertType(ossURL string) {
	switch setting.Conf.AlarmMedium.WebhookType {
	case "dingtalk":
		fmt.Println("dingtalk", "OSS URL", ossURL)
	case "email":
		fmt.Println("email", "OSS URL", ossURL)
	case "wechat":
		fmt.Println("wechat", "OSS URL", ossURL)
	default:
		fmt.Errorf("不支持该告警类型")
	}
}
