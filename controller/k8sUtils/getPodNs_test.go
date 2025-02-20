package k8sUtils

import (
	"heapdump_watcher/setting"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetPodNamespace(t *testing.T) {
	// 加载配置文件
	setting.InitConf()
	// k8s cient-go  [测试用的]
	clientset, err := setting.ReadKubeConf()
	if err != nil {
		logrus.Printf("ReadKubeConf Error: %s", err)
	}
	GetPodNamespace(clientset, "qi-capability-68896b6858-zwmfb")
}
