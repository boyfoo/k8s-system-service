package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"k8sapi/src/services"
)

//@controller
type NodeCtl struct {
	NodeService *services.NodeService `inject:"-"`
}

func NewNodeCtl() *NodeCtl {
	return &NodeCtl{}
}
func (this *NodeCtl) ListAll(c *gin.Context) goft.Json {
	return gin.H{
		"code": 20000,
		"data": this.NodeService.ListAllNodes(),
	}

}
func (this *NodeCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/nodes", this.ListAll)
}
func (*NodeCtl) Name() string {
	return "NodeCtl"
}
