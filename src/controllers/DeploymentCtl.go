package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
	"k8sapi/src/core"
)

type DeploymentCtl struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
	DepMap    *core.DeploymentMap   `inject:"-"` //注入， 首字母一定要大写
}

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}
func (this *DeploymentCtl) GetList(c *gin.Context) goft.Json {
	list, err := this.DepMap.ListByNS("default")
	goft.Error(err)
	return list
}
func (this *DeploymentCtl) Build(goft *goft.Goft) {
	//路由
	goft.Handle("GET", "/deployments", this.GetList)
}
func (*DeploymentCtl) Name() string {
	return "DeploymentCtl"
}
