package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type RBACCtl struct {
	RoleService *RoleService `inject:"-"`
}

func NewRBACCtl() *RBACCtl {
	return &RBACCtl{}
}
func (this *RBACCtl) Roles(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	return gin.H{
		"code": 20000,
		"data": this.RoleService.ListRoles(ns),
	}
}
func (*RBACCtl) Name() string {
	return "RBACCtl"
}
func (this *RBACCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/roles", this.Roles)
}
