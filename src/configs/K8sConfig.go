package configs

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8sapi/src/services"
	"log"
	"os/user"
)

type K8sConfig struct {
	DepHandler       *services.DepHandler       `inject:"-"`
	PodHandler       *services.PodHandler       `inject:"-"`
	NsHandler        *services.NsHandler        `inject:"-"`
	EventHandler     *services.EventHandler     `inject:"-"`
	IngressHandler   *services.IngressHandler   `inject:"-"`
	ServiceHandler   *services.ServiceHandler   `inject:"-"`
	SecretHandler    *services.SecretHandler    `inject:"-"`
	ConfigMapHandler *services.ConfigMapHandler `inject:"-"`
	NodeHandler      *services.NodeMapHandler   `inject:"-"`
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

//func(*K8sConfig) K8sRestConfig() *rest.Config{
//	config, err := clientcmd.BuildConfigFromFlags("","config" )
//	config.Insecure=true
//	if err!=nil{
//		log.Fatal(err)
//	}
//	return config
//}
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

//初始化Informer
func (this *K8sConfig) InitInformer() informers.SharedInformerFactory {
	fact := informers.NewSharedInformerFactory(this.InitClient(), 0)

	depInformer := fact.Apps().V1().Deployments()
	depInformer.Informer().AddEventHandler(this.DepHandler)

	podInformer := fact.Core().V1().Pods() //监听pod
	podInformer.Informer().AddEventHandler(this.PodHandler)

	serviceInformer := fact.Core().V1().Services() //监听service
	serviceInformer.Informer().AddEventHandler(this.ServiceHandler)

	nsInformer := fact.Core().V1().Namespaces() //监听namespace
	nsInformer.Informer().AddEventHandler(this.NsHandler)

	eventInformer := fact.Core().V1().Events() //监听event
	eventInformer.Informer().AddEventHandler(this.EventHandler)

	IngressInformer := fact.Networking().V1beta1().Ingresses() //监听Ingress
	IngressInformer.Informer().AddEventHandler(this.IngressHandler)

	SecretInformer := fact.Core().V1().Secrets() //监听Secret
	SecretInformer.Informer().AddEventHandler(this.SecretHandler)

	ConfigMapInformer := fact.Core().V1().ConfigMaps() //监听Configmap
	ConfigMapInformer.Informer().AddEventHandler(this.ConfigMapHandler)

	NodeInformer := fact.Core().V1().Nodes()
	NodeInformer.Informer().AddEventHandler(this.NodeHandler)

	fact.Start(wait.NeverStop)
	return fact
}
