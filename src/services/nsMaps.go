package services

import (
	corev1 "k8s.io/api/core/v1"
	"k8sapi/src/models"
	"sync"
)

// namespace相关
type NsMapStruct struct {
	data sync.Map // [key string] []*corev1.Namespace    key=>namespace的名称
}

func (this *NsMapStruct) Get(ns string) *corev1.Namespace {
	if item, ok := this.data.Load(ns); ok {
		return item.(*corev1.Namespace)
	}
	return nil
}
func (this *NsMapStruct) Add(ns *corev1.Namespace) {
	this.data.Store(ns.Name, ns)
}
func (this *NsMapStruct) Update(ns *corev1.Namespace) {
	this.data.Store(ns.Name, ns)
}
func (this *NsMapStruct) Delete(ns *corev1.Namespace) {
	this.data.Delete(ns.Name)
}

//显示所有的 namespace
func (this *NsMapStruct) ListAll() []*models.NsModel {
	ret := make([]*models.NsModel, 0)
	this.data.Range(func(key, value interface{}) bool {
		ret = append(ret, &models.NsModel{Name: key.(string)})
		return true
	})
	return ret
}
