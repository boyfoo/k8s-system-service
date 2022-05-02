package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type RBACCtl struct {
	RoleService *RoleService          `inject:"-"`
	Client      *kubernetes.Clientset `inject:"-"`
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
func (this *RBACCtl) RoleBindingList(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	return gin.H{
		"code": 20000,
		"data": this.RoleService.ListRoleBindings(ns),
	}
}
func (this *RBACCtl) DeleteRole(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	name := c.DefaultQuery("name", "")
	err := this.Client.RbacV1().Roles(ns).Delete(c, name, metav1.DeleteOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}
func (this *RBACCtl) CreateRole(c *gin.Context) goft.Json {
	role := rbacv1.Role{} //原生的k8s role 对象
	goft.Error(c.ShouldBindJSON(&role))
	role.APIVersion = "rbac.authorization.k8s.io/v1"
	role.Kind = "Role"
	_, err := this.Client.RbacV1().Roles(role.Namespace).Create(c, &role, metav1.CreateOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

func (*RBACCtl) Name() string {
	return "RBACCtl"
}
func (this *RBACCtl) Build(goft *goft.Goft) {
	goft.Handle("GET", "/roles", this.Roles)
	goft.Handle("GET", "/rolebindings", this.RoleBindingList)
	goft.Handle("POST", "/roles", this.CreateRole)
	goft.Handle("DELETE", "/roles", this.DeleteRole)
}
