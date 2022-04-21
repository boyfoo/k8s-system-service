package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

type DeploymentCtl struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
}

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}

func (d *DeploymentCtl) GetList(c *gin.Context) goft.Json {
	list, err := d.K8sClient.AppsV1().Deployments("default").List(c, metav1.ListOptions{})
	goft.Error(err)
	return list
}

func (d *DeploymentCtl) Build(goft *goft.Goft) {
	goft.Handle(http.MethodGet, "/deployments", d.GetList)
}

func (d *DeploymentCtl) Name() string {
	return "DeploymentCtl"
}
