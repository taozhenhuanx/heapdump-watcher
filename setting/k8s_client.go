package setting

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

/*
解决始终指定三个文件在go.mod文件中，换其他版本也可以,具体到github上看client-go的配套k8s版本
require (
    ...
    k8s.io/api v0.19.0
    k8s.io/apimachinery v0.19.0
    k8s.io/client-go v0.19.0
    ...
)
*/

func ReadKubeConf() (clientset *kubernetes.Clientset, err error) {
	var (
		config     *rest.Config // 可以在集群内部访问,也可以在集群外部访问。集群内部是在Pod中访问
		kubeconfig *string      // 集群外部是通过KubeConfig访问(Kubelet)
	)
	fmt.Println(Conf.FilePath.KubeConf)
	// 获取kubeconfig配置文件------这种方式是通过kubelet类型形式获取clientset
	if home := homeDir(); home != "" {
		// 如果home不等于空,在Linux上应该是等于/root,那么就拼接一个全路径 /root/.kube/config
		defaultKubeconfig := filepath.Join(home, ".kube", "config")
		if _, err := os.Stat(defaultKubeconfig); err == nil {
			// 如果默认路径的kubeconfig文件存在，则使用默认路径
			kubeconfig = &defaultKubeconfig
		} else if Conf.FilePath.KubeConf != "" {
			// 如果默认路径不存在，则使用从Conf.FilePath.KubeConf传入的路径
			kubeconfig = &Conf.FilePath.KubeConf
		} else {
			// 如果两者都不存在，则返回错误
			return nil, fmt.Errorf("默认路径的kubeconfig文件未找到")
		}
	}

	// 使用集群内部模式(ServiceAccount)获取clientset
	if config, err = rest.InClusterConfig(); err != nil {
		// 如果使用内部模式出错则使用外部模式 kubeconfig文件来创建集群配置
		if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
			return nil, fmt.Errorf("无法创建集群配置: %v", err)
		}
	}

	// 创建 Clientset 对象
	if clientset, err = kubernetes.NewForConfig(config); err != nil {
		return nil, fmt.Errorf("无法创建Clientset: %v", err)
	}

	return clientset, nil
}

// 获取kubeconfig文件的父路径
func homeDir() string {
	// 获取环境linux变量, HOME=/root
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	// 如果是Windows
	return os.Getenv("USERPROFILE")
}
