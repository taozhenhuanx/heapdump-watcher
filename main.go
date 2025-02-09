package main

import (
	"heapdump_watcher/controller/store/cli"
	"heapdump_watcher/setting"
)

func main() {
	// 加载配置文件
	setting.InitConf()
	// logrus.Infoln("渲染到结构体的参数：", setting.Conf)
	// 文件上传
	cli.UPload("x", "x")
}
