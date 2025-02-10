package dingtalk

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SendDingTalk() {
	url := fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", setting.Conf.AlarmMedium.DingTalkToken)

	payload := strings.NewReader(`{
    "msgtype": "link", 
    "link": {
        "text": "这个即将发布的新版本, 创始人xx称它为红树林。而在此之前, 每当面临重大升级, 产品经理们都会取一个应景的代号, 这一次, 为什么是红树林", 
        "title": "时代的火车向前开", 
        "picUrl": "", 
        "messageUrl": "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng=="
    }
}`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
