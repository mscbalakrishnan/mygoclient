package k8sclient

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetAllPods() {
	kubeconfig := flag.String("kubeconfig", "/Users/bnainar/.kube/config", "kube config file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		fmt.Println("error")
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("error in config")
	}

	podList, _ := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})

	for index, pod := range podList.Items {
		fmt.Println("index %i, name: {} ", index, pod.Name, pod.Status.Phase)
	}
}
