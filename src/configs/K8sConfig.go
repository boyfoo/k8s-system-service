package configs

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
)

type K8sConfig struct {
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

func (k *K8sConfig) InitClient() *kubernetes.Clientset {
	client, err := kubernetes.NewForConfig(&rest.Config{
		Host: "https://192.168.99.101:8443/",
		//
		TLSClientConfig: rest.TLSClientConfig{
			CAFile:   "/Users/rxt/.minikube/ca.crt",
			KeyFile:  "/Users/rxt/.minikube/profiles/minikube/client.key",
			CertFile: "/Users/rxt/.minikube/profiles/minikube/client.crt",
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	return client
}
