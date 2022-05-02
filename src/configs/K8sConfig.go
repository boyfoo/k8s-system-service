package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8sapi/pkg/rbac"
	"k8sapi/src/models"
	"k8sapi/src/services"
	"log"
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
	RoleHander       *rbac.RoleHander           `inject:"-"`
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

//初始化 系统 配置
func (*K8sConfig) InitSysConfig() *models.SysConfig {
	b, err := ioutil.ReadFile("app.yaml")
	if err != nil {
		log.Fatal(err)
	}
	config := &models.SysConfig{}
	err = yaml.Unmarshal(b, config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

//func(*K8sConfig) K8sRestConfig() *rest.Config{
//	config, err := clientcmd.BuildConfigFromFlags("","config" )
//	config.Insecure=true
//	if err!=nil{
//		log.Fatal(err)
//	}
//	return config
//}

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

	RolesInformer := fact.Rbac().V1().Roles()
	RolesInformer.Informer().AddEventHandler(this.RoleHander)

	fact.Start(wait.NeverStop)
	return fact
}
