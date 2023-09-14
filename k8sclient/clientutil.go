package k8sclient

import (
	"flag"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig *string

func init() {
	fmt.Print("Init method of k8sclient called.")
	if flag.Lookup("kubeconfig") == nil {
		kubeconfig = flag.String("kubeconfig", "/Users/bnainar/.kube/config", "kube config file")
	}
}

func getK8sClient() (*kubernetes.Clientset, error) {

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		fmt.Println("error")
	}
	return kubernetes.NewForConfig(config)
}
