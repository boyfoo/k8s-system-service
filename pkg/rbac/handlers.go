package rbac

import (
	"github.com/gin-gonic/gin"
	"k8s.io/api/rbac/v1"
	"k8sapi/src/wscore"
	"log"
)

type RoleHander struct {
	RoleMap     *RoleMapStruct `inject:"-"`
	RoleService *RoleService   `inject:"-"`
}

func (this *RoleHander) OnAdd(obj interface{}) {
	this.RoleMap.Add(obj.(*v1.Role))
	ns := obj.(*v1.Role).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "role",
			"result": gin.H{"ns": ns,
				"data": this.RoleService.ListRoles(ns)},
		},
	)
}
func (this *RoleHander) OnUpdate(oldObj, newObj interface{}) {
	err := this.RoleMap.Update(newObj.(*v1.Role))
	if err != nil {
		log.Println(err)
		return
	}
	ns := newObj.(*v1.Role).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "role",
			"result": gin.H{"ns": ns,
				"data": this.RoleService.ListRoles(ns)},
		},
	)
}
func (this *RoleHander) OnDelete(obj interface{}) {
	this.RoleMap.Delete(obj.(*v1.Role))
	ns := obj.(*v1.Role).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "role",
			"result": gin.H{"ns": ns,
				"data": this.RoleService.ListRoles(ns)},
		},
	)
}

type RoleBindingHander struct {
	RoleBindingMap *RoleBindingMapStruct `inject:"-"`
	RoleService    *RoleService          `inject:"-"`
}

func (this *RoleBindingHander) OnAdd(obj interface{}) {
	this.RoleBindingMap.Add(obj.(*v1.RoleBinding))
	ns := obj.(*v1.RoleBinding).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "rolebinding",
			"result": gin.H{"ns": ns,
				"data": this.RoleService.ListRoles(ns)},
		},
	)
}
func (this *RoleBindingHander) OnUpdate(oldObj, newObj interface{}) {
	err := this.RoleBindingMap.Update(newObj.(*v1.RoleBinding))
	if err != nil {
		log.Println(err)
		return
	}
	ns := newObj.(*v1.RoleBinding).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "rolebinding",
			"result": gin.H{"ns": ns,
				"data": this.RoleService.ListRoles(ns)},
		},
	)
}
func (this *RoleBindingHander) OnDelete(obj interface{}) {
	this.RoleBindingMap.Delete(obj.(*v1.RoleBinding))
	ns := obj.(*v1.Role).Namespace
	wscore.ClientMap.SendAll(
		gin.H{
			"type": "rolebinding",
			"result": gin.H{"ns": ns,
				"data": this.RoleService.ListRoles(ns)},
		},
	)
}
