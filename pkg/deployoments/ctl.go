package deployoments

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
	"k8sapi/src/services"
)

type DeploymentCtlV2 struct {
	K8sClient *kubernetes.Clientset   `inject:"-"`
	DeployMap *services.DeploymentMap `inject:"-"`
}

func NewDeploymentCtlV2() *DeploymentCtlV2 {
	return &DeploymentCtlV2{}
}
func (this *DeploymentCtlV2) LoadDeploy(c *gin.Context) goft.Json {
	ns := c.Param("ns")
	name := c.Param("name")
	dep, err := this.DeployMap.GetDeployment(ns, name) // 原生
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": dep,
	}
}

func (this *DeploymentCtlV2) Build(goft *goft.Goft) {
	//路由
	goft.Handle("GET", "/deployments/:ns/:name", this.LoadDeploy)
}
func (*DeploymentCtlV2) Name() string {
	return "DeploymentCtlV2"
}
