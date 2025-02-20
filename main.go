package main

import (
	"heapdump_watcher/controller/watchFile"
	"heapdump_watcher/setting"
)

func main() {
	// 加载配置文件
	setting.InitConf()
	// logrus.Infoln("渲染到结构体的参数：", setting.Conf)

	// 启动监听
	watchFile.WatchFiles()
}
