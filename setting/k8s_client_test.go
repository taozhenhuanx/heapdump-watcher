package setting

import (
	"context"
	"fmt"
	"log"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestReadKubeConf(t *testing.T) {
	clientset, err := ReadKubeConf()
	if err != nil {
		log.Printf("ReadKubeConf Error: %s", err)
	}
	fmt.Println("ReadKubeConf OK")
	// 使用 clientsent 获取 Deployments
	deployments, err := clientset.AppsV1().Deployments("kube-system").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for idx, deploy := range deployments.Items {
		log.Printf("%d -> %s\n", idx+1, deploy.Name)
	}
}
