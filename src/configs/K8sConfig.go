package configs

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8sapi/src/core"
	"log"
)

type K8sConfig struct {
	DepHandler *core.DepHandler `inject:"-"`
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

//初始化客户端
func (*K8sConfig) InitClient() *kubernetes.Clientset {
	config := &rest.Config{
		Host: "https://192.168.99.101:8443/",
		//
		TLSClientConfig: rest.TLSClientConfig{
			CAFile:   "/Users/rxt/.minikube/ca.crt",
			KeyFile:  "/Users/rxt/.minikube/profiles/minikube/client.key",
			CertFile: "/Users/rxt/.minikube/profiles/minikube/client.crt",
		},
	}
	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

//初始化Informer
func (this *K8sConfig) InitInformer() informers.SharedInformerFactory {
	fact := informers.NewSharedInformerFactory(this.InitClient(), 0)

	depInformer := fact.Apps().V1().Deployments()
	depInformer.Informer().AddEventHandler(this.DepHandler)

	fact.Start(wait.NeverStop)

	return fact
}
