package main

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	coreapiv1 "k8s.io/api/core/v1"
	resourcev1 "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

const kubeConfigFilePath = "./conf/k8s-config"

type K8sConfig struct {
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

func (k *K8sConfig) K8sRestConfig() *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigFilePath)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
func (k *K8sConfig) InitClient() *kubernetes.Clientset {
	c, err := kubernetes.NewForConfig(k.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}
	return c
}
func (k *K8sConfig) InitDynamicClient() dynamic.Interface {
	c, err := dynamic.NewForConfig(k.K8sRestConfig())
	if err != nil {
		log.Fatal(err)
	}
	return c
}
func (k *K8sConfig) InitDiscoveryClient() *discovery.DiscoveryClient {
	return discovery.NewDiscoveryClient(k.InitClient().RESTClient())
}

func CreateDeploy() {
	client := NewK8sConfig().InitClient()
	deploy := &v1.Deployment{}
	deploy.Namespace = "devops-test1"
	deploy.Name = "nginx-1-pre"
	deploy.Labels = map[string]string{
		"app":         "nginx-1",
		"env":         "pre",
		"app_service": "nginx-1-pre",
	}
	deploy.Annotations = map[string]string{
		"desc": "this is nginx-1 env for pre",
	}
	deploy.Spec.Replicas = int32Ptr(3)
	deploy.Spec.Selector = &metav1.LabelSelector{
		MatchLabels: map[string]string{
			"app_service": "nginx-1-pre",
		},
	}

	deploy.Spec.Template.Labels = map[string]string{
		"app_service": "nginx-1-pre",
	}

	container := coreapiv1.Container{
		Name:  "nginx-1-pre",
		Image: "nginx",
		Ports: []coreapiv1.ContainerPort{{
			ContainerPort: 80,
		}},
		Resources: coreapiv1.ResourceRequirements{
			Requests: coreapiv1.ResourceList{
				coreapiv1.ResourceCPU:    resourceMustParse("250m"),
				coreapiv1.ResourceMemory: resourceMustParse("64Mi"),
			},
			Limits: coreapiv1.ResourceList{
				coreapiv1.ResourceCPU:    resourceMustParse("500m"),
				coreapiv1.ResourceMemory: resourceMustParse("128Mi"),
			},
		},
		LivenessProbe:  &coreapiv1.Probe{},
		ReadinessProbe: &coreapiv1.Probe{},
		StartupProbe:   &coreapiv1.Probe{},
	}
	deploy.Spec.Template.Spec.Containers = []coreapiv1.Container{
		container,
	}

	opts := metav1.CreateOptions{}
	result, e := client.AppsV1().Deployments("devops-test1").Create(context.Background(), deploy, opts)
	if e != nil {
		fmt.Println("error: ", e.Error())
		return
	}
	bs, _ := json.Marshal(result)
	fmt.Println(string(bs))
}

func int32Ptr(i int32) *int32 { return &i }

// Helper function to parse resource quantities
func resourceMustParse(value string) resourcev1.Quantity {
	quantity, err := resourcev1.ParseQuantity(value)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse resource quantity: %v", err))
	}
	return quantity
}

func main() {
	CreateDeploy()
}
