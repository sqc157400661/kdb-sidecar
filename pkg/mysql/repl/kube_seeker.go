package repl

import "k8s.io/client-go/kubernetes"

type KubeSeeker struct {
	clientSet *kubernetes.Clientset
}

func NewKubeSeeker(clientSet *kubernetes.Clientset) *KubeSeeker {
	return &KubeSeeker{clientSet: clientSet}
}

func (s *KubeSeeker) GetHostInfoByHostname() {

}

func (s *KubeSeeker) GetHostInfoByClusterID() {

}
