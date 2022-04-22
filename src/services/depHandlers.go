package services

import (
	"k8s.io/api/apps/v1"
	"k8sapi/src/wscore"
	"log"
)

//处理deployment 回调的handler
type DepHandler struct {
	DepMap     *DeploymentMap     `inject:"-"`
	DepService *DeploymentService `inject:"-"`
}

func (this *DepHandler) OnAdd(obj interface{}) {
	this.DepMap.Add(obj.(*v1.Deployment))
	// 向所有的ws客户端发送
	wscore.ClientMap.SendAll(this.DepService.ListAll(obj.(*v1.Deployment).Namespace))
}
func (this *DepHandler) OnUpdate(oldObj, newObj interface{}) {
	err := this.DepMap.Update(newObj.(*v1.Deployment))
	if err != nil {
		log.Println(err)
	} else {
		wscore.ClientMap.SendAll(this.DepService.ListAll(newObj.(*v1.Deployment).Namespace))
	}
}
func (this *DepHandler) OnDelete(obj interface{}) {
	if d, ok := obj.(*v1.Deployment); ok {
		this.DepMap.Delete(d)
		wscore.ClientMap.SendAll(this.DepService.ListAll(obj.(*v1.Deployment).Namespace))
	}
}
