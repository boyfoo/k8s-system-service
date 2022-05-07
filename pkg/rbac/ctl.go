package rbac

import (
	"fmt"
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
func (this *RBACCtl) CreateRoleBinding(c *gin.Context) goft.Json {
	rb := &rbacv1.RoleBinding{}
	goft.Error(c.ShouldBindJSON(rb))
	_, err := this.Client.RbacV1().RoleBindings(rb.Namespace).Create(c, rb, metav1.CreateOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
	}
}

//从rolebinding中 增加或删除用户
func (this *RBACCtl) AddUserToRoleBinding(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	name := c.DefaultQuery("name", "") //rolebinding 名称
	t := c.DefaultQuery("type", "")    //如果没传值就是增加，传值（不管什么代表删除)
	subject := rbacv1.Subject{}        // 传过来
	goft.Error(c.ShouldBindJSON(&subject))
	if subject.Kind == "ServiceAccount" {
		subject.APIGroup = ""
	}
	rb := this.RoleService.GetRoleBinding(ns, name) //通过名称获取 rolebinding对象
	if t != "" {                                    //代表删除

		for i, sub := range rb.Subjects {
			if sub.Kind == subject.Kind && sub.Name == subject.Name {
				rb.Subjects = append(rb.Subjects[:i], rb.Subjects[i+1:]...)
				break //确保只删一个（哪怕有同名同kind用户)
			}
		}
		fmt.Println(rb.Subjects)
	} else {
		rb.Subjects = append(rb.Subjects, subject)
	}
	_, err := this.Client.RbacV1().RoleBindings(ns).Update(c, rb, metav1.UpdateOptions{})
	goft.Error(err)
	return gin.H{
		"code": 20000,
		"data": "success",
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
func (this *RBACCtl) DeleteRoleBinding(c *gin.Context) goft.Json {
	ns := c.DefaultQuery("ns", "default")
	name := c.DefaultQuery("name", "")
	err := this.Client.RbacV1().RoleBindings(ns).Delete(c, name, metav1.DeleteOptions{})
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
	goft.Handle("POST", "/rolebindings", this.CreateRoleBinding)
	goft.Handle("PUT", "/rolebindings", this.AddUserToRoleBinding) //添加用户到binding
	goft.Handle("DELETE", "/rolebindings", this.DeleteRoleBinding)
	goft.Handle("POST", "/roles", this.CreateRole)
	goft.Handle("DELETE", "/roles", this.DeleteRole)
}
