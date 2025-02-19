package wechat

import (
	"fmt"
	"heapdump_watcher/setting"
	"testing"
)

func TestSendWeChat(t *testing.T) {
	// 加载配置文件
	setting.InitConf()

	if err := SendWeChat("heapdump 告警信息,文件已经转存，请及时下载", "生产环境", "APP应用", "ossURL"); err != nil {
		fmt.Println(err)
	}
	fmt.Println("SendWeChat 成功")
}
