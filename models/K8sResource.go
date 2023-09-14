package models

import "fmt"

type K8sResource struct {
	Name         string `json:"name"`
	Status       string `json:"status"`
	ResourceType string `json:"resourceType"`
}

func (k8sResource K8sResource) ToString() {
	fmt.Printf("Name: %s, Status: %s, Resource Type: %s", k8sResource.Name, k8sResource.Status, k8sResource.ResourceType)
	fmt.Println("")
}
