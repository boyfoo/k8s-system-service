package rbac

import (
	"fmt"
	"k8s.io/api/rbac/v1"
	"sort"
	"sync"
)

type V1Role []*v1.Role

func (this V1Role) Len() int {
	return len(this)
}
func (this V1Role) Less(i, j int) bool {
	//根据时间排序    倒排序
	return this[i].CreationTimestamp.Time.After(this[j].CreationTimestamp.Time)
}
func (this V1Role) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type RoleMapStruct struct {
	data sync.Map // [ns string] []*v1.Role
}

func (this *RoleMapStruct) Get(ns string, name string) *v1.Role {
	if items, ok := this.data.Load(ns); ok {
		for _, item := range items.([]*v1.Role) {
			if item.Name == name {
				return item
			}
		}
	}
	return nil
}
func (this *RoleMapStruct) Add(item *v1.Role) {
	if list, ok := this.data.Load(item.Namespace); ok {
		list = append(list.([]*v1.Role), item)
		this.data.Store(item.Namespace, list)
	} else {
		this.data.Store(item.Namespace, []*v1.Role{item})
	}
}
func (this *RoleMapStruct) Update(item *v1.Role) error {
	if list, ok := this.data.Load(item.Namespace); ok {
		for i, range_item := range list.([]*v1.Role) {
			if range_item.Name == item.Name {
				list.([]*v1.Role)[i] = item
			}
		}
		return nil
	}
	return fmt.Errorf("Role-%s not found", item.Name)
}
func (this *RoleMapStruct) Delete(svc *v1.Role) {
	if list, ok := this.data.Load(svc.Namespace); ok {
		for i, range_item := range list.([]*v1.Role) {
			if range_item.Name == svc.Name {
				newList := append(list.([]*v1.Role)[:i], list.([]*v1.Role)[i+1:]...)
				this.data.Store(svc.Namespace, newList)
				break
			}
		}
	}
}
func (this *RoleMapStruct) ListAll(ns string) []*v1.Role {
	if list, ok := this.data.Load(ns); ok {
		newList := list.([]*v1.Role)
		sort.Sort(V1Role(newList)) //  按时间倒排序
		return newList
	}
	return []*v1.Role{} //返回空列表
}
