package internal

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GetClientSet() (cli *kubernetes.Clientset, err error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// 创建 Kubernetes 客户端
	return kubernetes.NewForConfig(config)
}
