package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"heapdump_watcher/setting"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// 发送信息的数据格式
type SendMsg struct {
	Msgtype  string    `json:"msgtype"`
	Markdown *Markdown `json:"markdown"`
}

type Markdown struct {
	Content string `json:"content"`
}

type Text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`        // userid的列表，提醒群中的指定成员(@某个成员)，@all表示提醒所有人
	MentionedMobileList []string `json:"mentioned_mobile_list"` // 手机号列表，提醒手机号对应的群成员(@某个成员)，@all表示提醒所有人
}

func SendWeChat(msg, env, appName, ossURL string) error {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", setting.Conf.AlarmMedium.WeChatKey)

	markdownText := fmt.Sprintf(
		"### 告警通知\n\n"+
			"> **环境**: %s\n\n"+
			"> **应用**: %s\n\n"+
			"> **OSS地址**: [点击查看](%s)\n\n"+
			"> **告警信息**: %s\n",
		env, appName, ossURL, msg)
	/*
		// 企业微信群机器人需要@的人
		mentioned_list := []string{}
		mentioned_mobile_list := []string{}
				sendMsg := SendMsg{
				Msgtype: "text",
				Text: &Text{
					Content:             content,
					MentionedList:       mentioned_list,
					MentionedMobileList: mentioned_mobile_list,
				},
			}
		content := fmt.Sprintf("JAVA 业务服务 OOM文件了, 请下载链接查看%s", ossURL)
	*/
	sendMsg := SendMsg{
		Msgtype: "markdown",
		Markdown: &Markdown{
			Content: markdownText,
		},
	}

	jsonData, err := json.Marshal(sendMsg)
	if err != nil {
		logrus.Println("WeChat JSON format conversion failed ", err)
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))

	if err != nil {
		logrus.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		logrus.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Println(err)
		return err
	}
	logrus.Println("企业微信通知发送成功!")
	logrus.Println(string(body))
	return nil
}
