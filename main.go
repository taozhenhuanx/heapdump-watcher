package main

import (
	"heapdump_watcher/controller/collectors"
	"heapdump_watcher/controller/watchFile"
	"heapdump_watcher/setting"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func main() {
	// 加载配置文件
	setting.InitConf()
	// logrus.Infoln("渲染到结构体的参数：", setting.Conf)

	// 创建一个监听器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logrus.Fatal("NewWatcher failed: ", err)
	}
	defer watcher.Close()

	// 启动监听, 传递watcher
	watchFile.WatchFiles(watcher)

	// 注册指标 + 定义指标
	prometheus.MustRegister(collectors.NewAlertCountCollector())

	// 注册控制器
	http.Handle("/metrics", promhttp.Handler())
	// 启动web服务
	if err := http.ListenAndServe(":9889", nil); err != nil {
		logrus.Fatal("ListenAndServe failed: ", err)
	}

	// 设置信号处理，等待程序退出
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// 等待信号
	<-sigs
}
