package dingtalk

import (
	"fmt"
	"heapdump_watcher/setting"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func SendDingTalk(msg, env, appName, ossURL string) error {
	webHook := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", setting.Conf.AlarmMedium.DingTalkToken)

	// 构造告警标题和 Markdown 内容
	title := fmt.Sprintf("【告警】%s", appName)
	markdownText := fmt.Sprintf(
		"### 告警通知\n\n"+
			"> **环境**: %s\n\n"+
			"> **应用**: %s\n\n"+
			"> **OSS地址**: [点击查看](%s)\n\n"+
			"> **告警信息**: %s\n",
		env, appName, ossURL, msg)

	// 构造 JSON 格式的消息体（使用 Markdown 格式）
	content := fmt.Sprintf(`{
	"msgtype": "markdown",
	"markdown": {
		"title": "%s",
		"text": "%s"
	},
	"at": {
		"atMobiles": [
			""
		],
		"isAtAll": false
	}
}`, title, markdownText)

	// 创建 POST 请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		logrus.Println("创建请求错误: ", err)
		return err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Println("发送请求错误: ", err)
		return err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Println("读取响应错误: ", err)
		return err
	}
	logrus.Println("响应内容: ", string(body))
	return nil
}
