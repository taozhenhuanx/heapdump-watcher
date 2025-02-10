package wechat

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SendWeChat() {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", setting.Conf.AlarmMedium.WeChatKey)
	payload := strings.NewReader(`{
    "msgtype": "text",
    "text": {
        "content": "今日天气：29度，大部分多云，降雨概率：60%",
		"mentioned_list":["wangqing","@all"],
		"mentioned_mobile_list":["13800001111","@all"]
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
