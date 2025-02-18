package main

import (
	"heapdump_watcher/controller/watchFile"
	"heapdump_watcher/setting"

	"github.com/sirupsen/logrus"
)

func main() {
	// 加载配置文件
	setting.InitConf()
	// logrus.Infoln("渲染到结构体的参数：", setting.Conf)

	// k8s cient-go  [测试用的]
	// clientset, err := setting.ReadKubeConf()
	// if err != nil {
	// 	log.Printf("ReadKubeConf Error: %s", err)
	// }
	// fmt.Println("ReadKubeConf OK")
	// // 使用 clientsent 获取 Deployments
	// deployments, err := clientset.AppsV1().Deployments("kube-system").List(context.TODO(), metav1.ListOptions{})
	// if err != nil {
	// 	panic(err)
	// }
	// for idx, deploy := range deployments.Items {
	// 	log.Printf("%d -> %s\n", idx+1, deploy.Name)
	// }
	watchFile.WatchFiles()
	logrus.Println("heapdump-watcher 程序已经启动")
}
