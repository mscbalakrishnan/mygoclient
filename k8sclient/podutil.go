package k8sclient

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mscbalakrishnan/mygoclient/models"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var k8sResourceList []models.K8sResource

type Cluster struct {
	Namespace string
}

func NewCluster(namespace string) Cluster {
	if namespace == "" {
		return Cluster{}
	} else {
		return Cluster{Namespace: namespace}
	}

}

func (c *Cluster) GetAllPods() []models.K8sResource {

	k8sResourceList = []models.K8sResource{}

	namespace := c.Namespace

	fmt.Printf("Namespace : %s", namespace)

	clientset, _ := getK8sClient()

	podList, _ := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})

	for index, pod := range podList.Items {
		fmt.Printf("index %ds, name: %s, status: %s ", index, pod.Name, pod.Status.Phase)
		k8sResource := models.K8sResource{}
		k8sResource.Name = string(pod.Name)
		k8sResource.Status = string(pod.Status.Phase)
		k8sResource.ResourceType = "POD"
		k8sResource.ToString()

		k8sResourceList = append(k8sResourceList, k8sResource)

	}

	return k8sResourceList
}

func (c *Cluster) GetAllDeployments() []models.K8sResource {

	k8sResourceList = []models.K8sResource{}

	clientset, _ := getK8sClient()
	deploymentList, err := clientset.AppsV1().Deployments(c.Namespace).List(context.Background(), metav1.ListOptions{})

	if err != nil {
		fmt.Printf("error in retrieving deployments.")
	}

	for _, deployment := range deploymentList.Items {

		k8sResource := models.K8sResource{}
		k8sResource.Name = string(deployment.Name)
		k8sResource.Status = strconv.Itoa(int(deployment.Status.AvailableReplicas)) + "/" + strconv.Itoa(int(deployment.Status.UpdatedReplicas))
		k8sResource.ResourceType = "DEPLOY"
		k8sResource.ToString()

		k8sResourceList = append(k8sResourceList, k8sResource)
	}

	return k8sResourceList
}
