package services

import (
	corev1 "k8s.io/api/core/v1"
)

// namespace 相关的回调handler
type NsHandler struct {
	NsMap *NsMapStruct `inject:"-"`
}

func (this *NsHandler) OnAdd(obj interface{}) {
	this.NsMap.Add(obj.(*corev1.Namespace))
}
func (this *NsHandler) OnUpdate(oldObj, newObj interface{}) {
	this.NsMap.Update(newObj.(*corev1.Namespace))

}
func (this *NsHandler) OnDelete(obj interface{}) {
	if d, ok := obj.(*corev1.Namespace); ok {
		this.NsMap.Delete(d)
	}
}
