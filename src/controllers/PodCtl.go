package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
	"k8sapi/src/services"
)

type PodCtl struct {
	K8sClient  *kubernetes.Clientset `inject:"-"`
	PodService *services.PodService  `inject:"-"` //首字母一定要大写
}

func NewPodCtl() *PodCtl {
	return &PodCtl{}
}
func (this *PodCtl) GetList(c *gin.Context) goft.Json {
	return this.PodService.ListByNs("istio-system")
}
func (this *PodCtl) Build(goft *goft.Goft) {
	//路由
	goft.Handle("GET", "/pods", this.GetList)
}
func (*PodCtl) Name() string {
	return "PodCtl"
}
