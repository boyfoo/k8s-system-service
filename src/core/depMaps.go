package core

import (
	"fmt"
	appv1 "k8s.io/api/apps/v1"
	"sync"
)

//对deployments的集合进行定义
type DeploymentMap struct {
	data sync.Map // [key string] []*v1.Deployment    key=>namespace
}

//添加
func (this *DeploymentMap) Add(dep *appv1.Deployment) {

	if list, ok := this.data.Load(dep.Namespace); ok {
		list = append(list.([]*appv1.Deployment), dep)
		this.data.Store(dep.Namespace, list)
	} else {
		this.data.Store(dep.Namespace, []*appv1.Deployment{dep})
	}
}

//更新
func (this *DeploymentMap) Update(dep *appv1.Deployment) error {
	if list, ok := this.data.Load(dep.Namespace); ok {
		for i, range_dep := range list.([]*appv1.Deployment) {
			if range_dep.Name == dep.Name {
				list.([]*appv1.Deployment)[i] = dep
			}
		}
		return nil
	}
	return fmt.Errorf("deployment-%s not found", dep.Name)
}

// 删除
func (this *DeploymentMap) Delete(dep *appv1.Deployment) {
	if list, ok := this.data.Load(dep.Namespace); ok {
		for i, range_dep := range list.([]*appv1.Deployment) {
			if range_dep.Name == dep.Name {
				newList := append(list.([]*appv1.Deployment)[:i], list.([]*appv1.Deployment)[i+1:]...)
				this.data.Store(dep.Namespace, newList)
				break
			}
		}
	}
}
func (this *DeploymentMap) ListByNS(ns string) ([]*appv1.Deployment, error) {
	if list, ok := this.data.Load(ns); ok {
		return list.([]*appv1.Deployment), nil
	}
	return nil, fmt.Errorf("record not found")
}
func (this *DeploymentMap) GetDeployment(ns string, depname string) (*appv1.Deployment, error) {
	if list, ok := this.data.Load(ns); ok {
		for _, item := range list.([]*appv1.Deployment) {
			if item.Name == depname {
				return item, nil
			}
		}
	}
	return nil, fmt.Errorf("record not found")
}
