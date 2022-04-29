package configs

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/metrics/pkg/client/clientset/versioned"
	"log"
	"os/user"
)

func (k *K8sConfig) InitClient() *kubernetes.Clientset {
	c, err := kubernetes.NewForConfig(k.InitConfig())
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func (*K8sConfig) InitConfig() *rest.Config {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	var config *rest.Config

	if "zx" == currentUser.Username {
		config = &rest.Config{
			Host:        "https://node01:8443/k8s/clusters/c-c9c5p",
			BearerToken: "kubeconfig-user-whq6g:ztjtlgcbs9vpqgcl9bqjpn598x629spvhdk74gv4kcpk8clvrddmzp",
			TLSClientConfig: rest.TLSClientConfig{ // 不验证响应tls
				Insecure: true,
			},
		}
	} else {
		config = &rest.Config{
			Host: "https://192.168.99.101:8443/",
			//
			TLSClientConfig: rest.TLSClientConfig{
				CAFile:   "/Users/rxt/.minikube/ca.crt",
				KeyFile:  "/Users/rxt/.minikube/profiles/minikube/client.key",
				CertFile: "/Users/rxt/.minikube/profiles/minikube/client.crt",
			},
		}
	}

	return config
}

// metric客户端
func (this *K8sConfig) InitMetricClient() *versioned.Clientset {
	c, err := versioned.NewForConfig(this.InitConfig())
	if err != nil {
		log.Fatal(err)
	}
	return c
}
