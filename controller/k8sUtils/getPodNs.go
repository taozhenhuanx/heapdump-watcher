package k8sUtils

import (
	"context"
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// getPodNamespace 根据 Pod 名称查找其命名空间
func GetPodNamespace(clientset *kubernetes.Clientset, podName string) (string, error) {
	// 获取所有命名空间
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to list namespaces: %v", err)
	}

	// 尝试在每个命名空间中寻找该 Pod
	for _, ns := range namespaces.Items {
		_, err = clientset.CoreV1().Pods(ns.Name).Get(context.TODO(), podName, metav1.GetOptions{})
		if err == nil {
			return ns.Name, nil // 找到 Pod，返回其命名空间
		}
		// 检查 Pod 不存在的情况
		if strings.Contains(err.Error(), "not found") {
			continue // Pod 不存在，继续检查下一个命名空间
		}
	}

	return "", fmt.Errorf("pod '%s' not found in any namespace", podName)
}
