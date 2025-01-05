package repl

import (
	"context"
	"fmt"
	"github.com/sqc157400661/kdb-sidecar/internal"
	"github.com/sqc157400661/kdb-sidecar/pkg/mysql/config"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strconv"
)

type KubeSeeker struct {
	clientSet *kubernetes.Clientset
}

func NewKubeSeeker(clientSet *kubernetes.Clientset) *KubeSeeker {
	return &KubeSeeker{clientSet: clientSet}
}

func (s *KubeSeeker) GetHostInfoByPodName(podName string) (host *HostInfo, err error) {
	pod, err := s.clientSet.CoreV1().Pods(config.K8SNamespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return
	}
	portStr := pod.Annotations[internal.MySQLPortAnno]
	var port = config.MySQLPort
	if portStr != "" {
		port, _ = strconv.Atoi(portStr)
	}
	host = &HostInfo{
		Host: pod.Status.PodIP,
		IP:   pod.Status.PodIP,
		Port: port,
	}
	return
}

func (s *KubeSeeker) GetHostInfoByClusterID(clusterID string) (hosts []*HostInfo, err error) {
	labelSelector := fmt.Sprintf("%s = %s", internal.MySQLClusterID, clusterID)
	pods, err := s.clientSet.CoreV1().Pods(config.K8SNamespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range pods.Items {
		portStr := pod.Annotations[internal.MySQLPortAnno]
		var port = config.MySQLPort
		if portStr != "" {
			port, _ = strconv.Atoi(portStr)
		}
		hosts = append(hosts, &HostInfo{
			Host: pod.Status.PodIP,
			IP:   pod.Status.PodIP,
			Port: port,
		})
	}
	return
}
